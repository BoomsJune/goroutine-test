package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func doOnce() {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("just once")
	})
}

func main() {
	wg.Add(3)

	go doOnce()
	go doOnce()
	go doOnce()

	wg.Wait()
}
