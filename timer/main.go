package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 5)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-timer.C
		fmt.Println("time's up!")
	}()

	fmt.Println("end")
	wg.Wait()
}
