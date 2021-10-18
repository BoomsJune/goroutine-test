package main

import (
	"fmt"
	"time"
)

func task(done chan string) {
	fmt.Println("doing somthing...")
	time.Sleep(time.Second * 3)
	done <- "Msg from task1"
}

func taskWatcher(done chan string) {
	select {
	case res := <-done:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout!")
	}
}

func main() {
	ch := make(chan string)
	go task(ch)
	taskWatcher(ch)
}
