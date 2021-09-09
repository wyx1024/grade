package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	tamp := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		defer close(tamp)
		fmt.Println("begin Wait")
		w.Wait()
		fmt.Println("end Wait")
	}()

	select {
	case <-tamp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	argumemts := os.Args
	if len(argumemts) != 2 {
		panic("not duration arge")
	}
	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi((argumemts[1]))
	if err != nil{
		panic(t)
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}


}
