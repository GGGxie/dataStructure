package main

import (
	"fmt"
	"time"
)

// 控制 goroutine 数量
type Pool struct {
	MaxNum   int
	TaskList chan int
}

func New(num int) *Pool {
	pool := &Pool{
		MaxNum:   num,
		TaskList: make(chan int),
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
	for task := range p.TaskList {
		fmt.Println(task)
	}
}

func (p *Pool) AddWork(task int) {
	p.TaskList <- task
}

func main() {
	p := New(10)
	p.AddWork(10)
	p.AddWork(11)
	p.AddWork(12)
	for {
		time.Sleep(time.Second)
	}
}
