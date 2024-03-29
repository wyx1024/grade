package balanceter

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIndex int
	rss      []string

	//conf LoadBalanceConf
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	add := params[0]
	r.rss = append(r.rss, add)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

//func (r *RandomBalance) SetConf(conf LoadBalanceConf) {
//	r.conf = conf
//}

func (r *RandomBalance) Update() {
	//if conf, ok := r.conf.(*LoadBalanceZkConf); ok {
	//	fmt.Println("Update get conf:", conf.GetConf())
	//	r.rss = []string{}
	//	for _, ip := range conf.GetConf() {
	//		r.Add(strings.Split(ip, ",")...)
	//	}
	//}
	//if conf, ok := r.conf.(*LoadBalanceCheckConf); ok {
	//	fmt.Println("Update get conf:", conf.GetConf())
	//	r.rss = nil
	//	for _, ip := range conf.GetConf() {
	//		r.Add(strings.Split(ip, ",")...)
	//	}
	//}
}
