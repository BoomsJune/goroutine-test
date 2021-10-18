package main

import (
	"fmt"
	"time"
)

func notifier(stop chan struct{}) {
	fmt.Println("notifier: wait 5 seconds")
	time.Sleep(time.Second * 5)
	fmt.Println("notifier: notify subscriber stop, then wait 5 seconds")
	stop <- struct{}{}
	time.Sleep(time.Second * 5)
	fmt.Println("notifier: done.")
}

func subscriber(stop chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("subscriber: stoped!")
			return
		default:
			fmt.Println("subscriber: running...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan struct{})
	go subscriber(ch)
	notifier(ch)
}
