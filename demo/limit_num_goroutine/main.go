package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func doSomething(u string) { // 模拟抓取任务的执行
	fmt.Println(u, time.Now().GoString())
	time.Sleep(2 * time.Second)
}

const (
	Limit  = 2 // 同時并行运行的goroutine上限
	Weight = 1 // 每个goroutine获取信号量资源的权重
)

func main() {
	urls := []string{
		"http://www.example.com",
		"http://www.example.net",
		"http://www.example.net/foo",
		"http://www.example.net/bar",
		"http://www.example.net/baz",
	}
	s := semaphore.NewWeighted(Limit)
	// w 防止最后一个在完成之前退出
	var w sync.WaitGroup
	for _, u := range urls {
		w.Add(1)
		s.Acquire(context.Background(), Weight)
		go func(u string) {
			doSomething(u)
			s.Release(Weight)
			w.Done()
		}(u)
	}
	w.Wait()

	fmt.Println("All Done")
}
