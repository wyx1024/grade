package main

import (
	"fmt"
	"go-growth/balanceter"
	"time"
)

func main() {
	selects := balanceter.LoadBalanceFactory(balanceter.LbConsistentHash)
	addrs := make( []string, 10)
	for i := 0; i < 10; i++ {
		addr := fmt.Sprintf("127.0.0.1:850%d", i)
		addrs = append(addrs, addr)
		//balance.Add(addr, strconv.Itoa(i+rand.Intn(10000)))
		selects.Add(addr)
	}
	for _, addr := range addrs {
		fmt.Println(selects.Get(addr))
		time.Sleep(time.Second)
	}
}