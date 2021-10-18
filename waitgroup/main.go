package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 计数器 = 2

	go func() {
		defer wg.Done() // 计数器 - 1
		fmt.Println("hello1")
	}()

	go func() {
		defer wg.Done() // 计数器 - 1
		fmt.Println("hello2")
	}()

	wg.Wait() // 等待计数器 = 0 终止
}
