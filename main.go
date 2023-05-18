package main

import (
	"errors"
	"fmt"
	"sync"
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

func main() {
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
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
