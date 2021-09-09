package main

import (
	"fmt"
	"go-growth/limiter"
	"time"
)

func main() {
	limit := limiter.NewBucketLimit(50, 50)
	m := make(map[int]bool)
	for i := 0; i < 153; i++ {
		allow := limit.Allow()
		if allow {
			continue
		}
		m[i] = false
		time.Sleep(time.Second)
	}
	fmt.Println(len(m))
	for i, b := range m {
		fmt.Println("i=",i, "b=",b)
	}

}
