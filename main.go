package main

import (
	"fmt"
	"sync"
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

func applyWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("hello1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("hello2")
	}()
	wg.Wait()
}

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

var once sync.Once

func doOnce() {
	defer wg.Done()
	once.Do(func() {
		fmt.Println("just once")
	})
}

func main() {
	// ch := make(chan int)
	// go producer(ch, 10)
	// consumer(ch)

	// ch := make(chan struct{})
	// go subscriber(ch)
	// notifier(ch)

	// ch := make(chan string)
	// go task(ch)
	// taskWatcher(ch)

	// applyWaitGroup()

	// wg.Add(3)
	// go add()
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(sum)

	// wg.Add(3)
	// go doOnce()
	// go doOnce()
	// go doOnce()
	// wg.Wait()
}
