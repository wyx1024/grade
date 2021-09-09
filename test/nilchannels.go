package main

import (
	"fmt"
	"math/rand"
	"time"
)
//一个 nil channel 总是阻塞的。
func Add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)
	for {
		select {
		case input := <-c:
			sum += input
			fmt.Println(input)
		case <-t.C:
			c = nil
			fmt.Println("sum:", sum)
		}
	}
}

func Send(c chan int) {
	for {
		n := rand.Intn(10)
		fmt.Println("send:", n)
		c <- n

	}
}

func main() {
	c := make(chan int)
	go Add(c)
	go Send(c)
	time.Sleep(time.Second * 3)
}
