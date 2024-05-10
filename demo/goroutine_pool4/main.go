package main

import (
	"fmt"
	"time"
)

// 动态控制 goroutine 数量
type Pool struct {
	MaxNum   chan int
	TaskList chan int
}

func New(num int) *Pool {
	pool := &Pool{
		MaxNum:   make(chan int, num),
		TaskList: make(chan int),
	}
	go pool.Start()
	return pool
}

func (p *Pool) Start() {
	for task := range p.TaskList {
		p.MaxNum <- 1
		go p.Work(task)
	}
}

func (p *Pool) Work(task int) {
	defer func() {
		<-p.MaxNum
	}()
	fmt.Println(task)
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
