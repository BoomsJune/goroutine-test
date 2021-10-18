package main

import (
	"fmt"
	"sync"
)

var sum int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock() // 加锁，其他goroutine等待
		sum += i
		lock.Unlock() // 解锁，等待的goroutine可进入
	}
}

func main() {
	wg.Add(3)

	go add()
	go add()
	go add()

	wg.Wait()
	fmt.Println(sum)
}
