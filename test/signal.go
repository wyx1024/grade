package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

func main() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign,os.Interrupt, syscall.SIGUSR1)//声明接收的信号类型
	go func() {
		for {
			sig :=<- sign
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught", sig)
			case syscall.SIGUSR1:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(time.Second*20)
	}
}