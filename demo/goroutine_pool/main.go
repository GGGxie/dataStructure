package main

import (
	"fmt"
)

type Worker struct { //协程池
	Ch  chan int
	Num int
	F   func(<-chan int, int)
}

func Constructor(num int, f func(<-chan int, int)) *Worker { //构造函数
	return &Worker{
		Ch:  make(chan int),
		Num: num,
		F:   f,
	}
}

func (w *Worker) Start() { //启动num个协程
	for i := 0; i < w.Num; i++ {
		go w.F(w.Ch, i)
	}
}

func (w *Worker) Add(i int) { //往协程池添加任务
	w.Ch <- i
}

func Work(ch <-chan int, idx int) { //goroutine执行的函数
	for i := range ch {
		fmt.Printf("worker %d : %d\n", idx, i)
	}
}

func main() {
	w := Constructor(2, Work)
	w.Start()
	for i := 0; i < 100; i++ {
		w.Add(i)
	}

	for{}
}
