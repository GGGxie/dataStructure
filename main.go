package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
	a          byte
)

type A struct {
	Str  string `json:"str,omitempty"`
	Str2 string `json:"str2,omitempty"`
}

func (a *A) Change() {
	a.Str = "3"
}
func (a A) Change2() {
	a.Str = "2"
}

var wg sync.WaitGroup

func close() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	select {
	case sig := <-sc:
		fmt.Println("Program Exit...", sig)
		// 各种回收
		os.Exit(1)
	}
	wg.Done()
}
func main() {
	fmt.Println(confusingNumber(11))
}

func confusingNumber(n int) bool {
	// 记录翻转后的数字
	mapp := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	// 位处理
	var changeNum int

	for tmp := n; tmp != 0; tmp = tmp / 10 {
		z := tmp % 10
		if c, ok := mapp[z]; ok {
			changeNum = changeNum*10 + c
		} else {
			return false
		}
	}
	return changeNum != n
}

type Array struct {
	data   []interface{}
	length int
}

func NewArray(size int) Array {
	return Array{
		data:   make([]interface{}, size),
		length: size,
	}
}

func (a Array) Get(idx int) (interface{}, error) {
	if idx > a.length-1 {
		return nil, errors.New("访问越界")
	}
	return a.data[idx], nil
}

func (a Array) Set(idx int, value interface{}) error {
	if idx > a.length-1 {
		return errors.New("访问越界")
	}
	a.data[idx] = value
	return nil
}
