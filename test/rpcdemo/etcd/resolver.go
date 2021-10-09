package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"log"

	"google.golang.org/grpc/resolver"
)

const schema = "etcd"

type MyResolver struct {
	endPoints []string
	service   string
	cli       *clientv3.Client
	cc        resolver.ClientConn
	addrDict  map[string]resolver.Address
}

func NewResolver(endpoints []string, service string) (resolver.Builder, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})

	if err != nil {
		return nil, err
	}

	return &MyResolver{
		endPoints: endpoints,
		cli:       client,
		service:   service}, nil
}

func (r *MyResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (r *MyResolver) Close() {}

func (r *MyResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	r.addrDict = make(map[string]resolver.Address)
	response, err := r.cli.Get(context.Background(), r.service, clientv3.WithPrefix())
	if err == nil {
		return nil, err
	}

	for _, kv := range response.Kvs {
		info := &ServiceInfo{}
		err := json.Unmarshal(kv.Value, info)
		if err != nil {
			log.Println(err)
		} else {
			r.addrDict[string(kv.Value)] = resolver.Address{Addr: info.IP}
		}
	}

	r.updateState()

	go r.watcher()

	return r, err
}

func (r *MyResolver) Scheme() string {
	return schema + "/" + r.service + "/"
}

func (r *MyResolver) watcher() {
	rch := r.cli.Watch(context.Background(), r.service, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for ev := range rch {
		update := false
		for _, event := range ev.Events {
			key := string(event.Kv.Value)
			switch event.Type {
			case mvccpb.PUT:
				info := &ServiceInfo{}
				err := json.Unmarshal(event.Kv.Value, info)
				if err != nil {
					log.Println(err)
				} else {
					r.addrDict[string(event.Kv.Value)] = resolver.Address{Addr: info.IP}
					update = true
				}
			case mvccpb.DELETE:
				_, ok := r.addrDict[key]
				if ok {
					delete(r.addrDict, key)
					update = true
				} else {
					log.Println(errors.New(fmt.Sprintf("not found Key %s", key)))
				}
			}
		}

		if update {
			r.updateState()
		}
	}
}

func (r *MyResolver) updateState() {
	if len(r.addrDict) == 0 {
		return
	}
	addrList := make([]resolver.Address, 0, len(r.addrDict))
	for _, v := range r.addrDict {
		addrList = append(addrList, v)
	}
	r.cc.UpdateState(resolver.State{
		Addresses: addrList,
	})
}
