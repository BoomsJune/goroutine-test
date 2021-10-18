package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 并发队列执行爬虫
// 1. 多个爬虫并发执行，单个爬虫循环执行
// 2.

type Job struct {
	ID int
}

var wg sync.WaitGroup
var parentCtx = context.Background()

func process(job Job) {
	ctx, cancel := context.WithTimeout(parentCtx, time.Second*2)
	defer cancel()
	defer wg.Done()

	done := make(chan struct{})
	go func() {
		if job.ID == 2 {
			time.Sleep(time.Second * 5)
		} else {
			time.Sleep(time.Second)
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Printf("Job %d done.\n", job.ID)
	case <-ctx.Done():
		fmt.Printf("Job %d timeout!\n", job.ID)
	}

}

func workerPool(size int, jobCh chan Job) {

	for i := 0; i < size; i++ {
		go worker(i, jobCh)
	}
}

func worker(num int, jobCh <-chan Job) {
	for job := range jobCh {
		fmt.Printf("worker %d: receive job %d \n", num, job.ID)
		process(job)
	}
}

func main() {
	jobCh := make(chan Job, 10)

	workerPool(3, jobCh)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		jobCh <- Job{ID: i + 1}
	}
	close(jobCh) // break range
	wg.Wait()
}
