package main

import (
	"fmt"
	"time"
)

type Pool struct {
	MaxNum int
	Task   chan int
}

func New(num int) *Pool {
	pool := &Pool{
		MaxNum: num,
		Task:   make(chan int),
	}
	go pool.Start()
	return pool
}

func (p *Pool) Start() {
	for i := 0; i < p.MaxNum; i++ {
		p.Work()
	}
}

func (p *Pool) Work() {
	for task := range p.Task {
		fmt.Println(task)
	}
}

func (p *Pool) AddWork(task int) {
	p.Task <- task
}

// type Pool struct {
// 	MaxNum chan int
// 	Task   chan int
// }

// func New(num int) *Pool {
// 	pool := &Pool{
// 		MaxNum: make(chan int, num),
// 		Task:   make(chan int),
// 	}
// 	go pool.Start()
// 	return pool
// }

// func (p *Pool) Start() {
// 	for task := range p.Task {
// 		p.MaxNum <- 1
// 		go p.Work(task)
// 	}
// }

// func (p *Pool) Work(task int) {
// 	defer func() {
// 		<-p.MaxNum
// 	}()
// 	fmt.Println(task)
// }

// func (p *Pool) AddWork(task int) {
// 	p.Task <- task
// }

func main() {
	p := New(10)
	p.AddWork(10)
	p.AddWork(11)
	p.AddWork(12)
	for {
		time.Sleep(time.Second)
	}
}

// func worker(tasks <-chan Task, results chan<- Result) {
// 	for task := range tasks {
// 		// 处理任务
// 		result := processTask(task)
// 		// 将结果发送到结果通道
// 		results <- result
// 	}
// }

// func main() {
// 	tasks := make(chan Task, 100)     // 任务通道
// 	results := make(chan Result, 100) // 结果通道

// 	// 启动固定数量的协程
// 	for i := 0; i < 10; i++ {
// 		go worker(tasks, results)
// 	}

// 	// 向任务通道发送任务
// 	for _, task := range tasksSlice {
// 		tasks <- task
// 	}
// 	close(tasks)

// 	// 从结果通道读取结果
// 	for i := 0; i < len(tasksSlice); i++ {
// 		result := <-results
// 		// 处理结果
// 	}
// }
