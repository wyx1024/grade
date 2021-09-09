package balances

type RoundRobinBalance struct {
	curIndex int
	rss      []string

	//conf LoadBalanceConf
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return nil
	}
	addr := params[0]
	r.rss = append(r.rss, addr)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	lens := len(r.rss)
	if lens == 0 {
		return ""
	}
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	addr := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return addr
}

func (r *RoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}

//func (r *RoundRobinBalance) SetConf(conf LoadBalanceConf) {
//	r.conf = conf
//}

func (r *RoundRobinBalance) Update() {
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
