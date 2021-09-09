package balanceter

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func([]byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ConsistentHashBalance struct {
	mux      sync.RWMutex
	hash     Hash
	replicas int               //复制因子 虚拟节点数
	keys     Uint32Slice       //已排序的节点hash切片 映射在环上的虚拟节点
	hashMap  map[uint32]string //节点哈希和Key的map,键是hash值，值是节点key

	//conf LoadBalanceConf
}

func NewConsistentHashBalance(replicas int, fn Hash) *ConsistentHashBalance {
	m := &ConsistentHashBalance{
		hash:     fn,
		replicas: replicas,
		hashMap: make(map[uint32]string, 100),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 验证是否为空
func (c *ConsistentHashBalance) IsEmpty() bool {
	return len(c.keys) == 0
}

func (c *ConsistentHashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	c.mux.Lock()
	defer c.mux.Unlock()
	for i := 0; i < c.replicas; i++ {
		hash := c.hash([]byte(strconv.Itoa(i) + addr))
		c.keys = append(c.keys, hash)
		c.hashMap[hash] = addr
	}
	sort.Sort(c.keys)
	return nil
}

func (c *ConsistentHashBalance) Get(key string) (string, error) {
	if c.IsEmpty() {
		return "", errors.New("node is empty")
	}

	hash := c.hash([]byte(key))
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})
	if idx == len(c.keys) {
		idx =0
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.hashMap[c.keys[idx]], nil

}

//func (c *ConsistentHashBalance) SetConf(conf LoadBalanceConf) {
//	c.conf = conf
//}

func (c *ConsistentHashBalance) Update() {
	//if conf, ok := c.conf.(*LoadBalanceZkConf); ok {
	//	fmt.Println("Update get conf:", conf.GetConf())
	//	c.mux.Lock()
	//	defer c.mux.Unlock()
	//	c.keys = nil
	//	c.hashMap = nil
	//	for _, ip := range conf.GetConf() {
	//		c.Add(strings.Split(ip, ",")...)
	//	}
	//}
	//if conf, ok := c.conf.(*LoadBalanceCheckConf); ok {
	//	fmt.Println("Update get conf:", conf.GetConf())
	//	c.mux.Lock()
	//	defer c.mux.Unlock()
	//	c.keys = nil
	//	c.hashMap = nil
	//	for _, ip := range conf.GetConf() {
	//		c.Add(strings.Split(ip, ",")...)
	//	}
	//}
}
