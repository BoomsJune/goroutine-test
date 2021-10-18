package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int, max int) {
	defer close(ch)
	for i := 0; i < max; i++ {
		fmt.Println("producer++:", i)
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("consumer--:", val)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch, 10)
	consumer(ch)
}
