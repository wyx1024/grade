package balanceter

import (
	"errors"
	"strconv"
)

type WeightNode struct {
	addr            string
	weight          int
	curWeight       int
	effectiveWeight int
}

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int

	//conf LoadBalanceConf
}

func (w *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	weight, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{
		addr:            params[0],
		weight:          int(weight),
		effectiveWeight: int(weight),
	}
	w.rss = append(w.rss, node)
	return nil
}

func (w *WeightRoundRobinBalance) Next() string {
	total := 0
	var bast *WeightNode
	for _, node := range w.rss {
		total += node.effectiveWeight
		node.curWeight += node.effectiveWeight
		if node.effectiveWeight < node.weight {
			node.effectiveWeight += 1
		}
		if bast == nil || node.curWeight > bast.curWeight {
			bast = node
		}
	}
	if bast == nil {
		return ""
	}
	bast.curWeight -= total
	return bast.addr
}

func (w *WeightRoundRobinBalance) Get(key string) (string, error) {
	return w.Next(), nil
}

//func (r *WeightRoundRobinBalance) SetConf(conf LoadBalanceConf) {
//	r.conf = conf
//}

func (w *WeightRoundRobinBalance) Update() {
	//if conf, ok := r.conf.(*LoadBalanceZkConf); ok {
	//	fmt.Println("WeightRoundRobinBalance get conf:", conf.GetConf())
	//	r.rss = nil
	//	for _, ip := range conf.GetConf() {
	//		r.Add(strings.Split(ip, ",")...)
	//	}
	//}
	//if conf, ok := r.conf.(*LoadBalanceCheckConf); ok {
	//	fmt.Println("WeightRoundRobinBalance get conf:", conf.GetConf())
	//	r.rss = nil
	//	for _, ip := range conf.GetConf() {
	//		r.Add(strings.Split(ip, ",")...)
	//	}
	//}
}
