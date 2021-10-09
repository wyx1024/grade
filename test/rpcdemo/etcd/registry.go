package etcd

import (
	"context"
	"encoding/json"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

type ServiceInfo struct {
	Name string
	IP   string
}

type Service struct {
	ServiceInfo ServiceInfo
	stop        chan error
	lessID      clientv3.LeaseID
	client      *clientv3.Client
}

func NewService(info ServiceInfo, endPoints []string) (*Service, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: time.Second * 20,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	s := &Service{
		ServiceInfo: info,
		client:      client,
	}
	return s, nil
}

func (s *Service) Start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Fatal(err)
		return err
	}
	for {
		select {
		case <-s.stop:
			return errors.New("service stop")
		case <-s.client.Ctx().Done():
			return errors.New("service Client Ctx Done")
		case _, ok := <-ch:
			if !ok {
				log.Println("keep alive channel closed")
				return s.revoke()
			}
			//log.Printf("Recv reply from service: %s, ttl:%d", s.getKey(), resp.TTL)
		}
	}
}

func (s *Service) Stop() {
	close(s.stop)

	s.stop <- nil
}

func (s *Service) revoke() error {
	if _, err := s.client.Revoke(context.Background(), s.lessID); err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("service %s stop/m", s.getKey)
	return nil
}

func (s *Service) getKey() string {
	return schema + "/" + s.ServiceInfo.Name + "/" + s.ServiceInfo.IP
}

func (s *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	key := s.getKey()
	val, _ := json.Marshal(s.ServiceInfo)
	grant, err := s.client.Grant(context.Background(), 5)
	if err != nil {
		return nil, err
	}
	if _, err = s.client.Put(context.Background(), key, string(val), clientv3.WithLease(grant.ID)); err != nil {
		return nil, err
	}
	s.lessID = grant.ID

	return s.client.KeepAlive(context.Background(), grant.ID)
}
