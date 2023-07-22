package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t := New(2)
	t.Go(
		func() {
			fmt.Println("1")
		},
	)
	t.Go(
		func() {
			time.Sleep(5 * time.Second)
			fmt.Println("1")
		},
	)
	t.Go(
		func() {
			fmt.Println("1")
		},
	)
	t.WaitGroup()
	time.Sleep(5 * time.Second)
}

type Task struct {
	Num    int
	sc     *sync.WaitGroup
	ch     chan func()
	cancel chan struct{}
}

func New(num int) *Task {
	ret := &Task{
		Num:    num,
		sc:     &sync.WaitGroup{},
		ch:     make(chan func()),
		cancel: make(chan struct{}),
	}
	for i := 0; i < num; i++ {
		go ret.Handle()
	}
	return ret
}
func (t *Task) Handle() {
	for {
		select {
		case f := <-t.ch:
			f()
			t.sc.Done()
		case <-t.cancel:
			fmt.Println("return ")
			return
		}
	}
}
func (t *Task) Go(f func()) {
	t.sc.Add(1)
	t.ch <- f
}
func (t *Task) WaitGroup() {
	t.sc.Wait()
	close(t.cancel)
}
