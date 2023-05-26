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
	fmt.Println(-10 % 3)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func stringShift2(s string, shift [][]int) string {
	// 计算移动结果
	var mov int
	for i := range shift {
		if shift[i][0] == 0 {
			mov += shift[i][1]
		} else {
			mov -= shift[i][1]
		}
	}
	length := len(s)
	var idx int
	if mov > 0 {
		idx = mov % length
	} else {
		idx = (mov % length) + length
	}
	return s[idx:] + s[:idx]
}

// 相隔为 1 的编辑距离
// https://leetcode.cn/problems/one-edit-distance/
func isOneEditDistance(s string, t string) bool {
	distance := len(s) - len(t)
	if distance == 1 && (len(s) == 0 || len(t) == 0) {
		return true
	}
	var count int
	switch distance { //distance有三种情况
	case 0:
		{ //长度相等
			for i := range s {
				if s[i] != t[i] {
					count++
				}
			}
			return count == 1
		}
	case 1:
		{ //s 比 t 多一个
			for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
				if s[i] != t[j] { //遇到不同的直接比较后半段
					count++
					return s[i+1:] == t[j:]
				}
			}
			return true //全部遍历完说明前面都相同,就最后一个字符不同,返回 true
		}
	case -1:
		{ //s 比 t 少一个
			for i, j := 0, 0; i < len(t) && j < len(s); i, j = i+1, j+1 {
				if s[j] != t[i] { //遇到不同的直接比较后半段
					count++
					return t[i+1:] == s[j:]
				}
			}
			return true //全部遍历完说明前面都相同,就最后一个字符不同,返回 true
		}
	default: //长度相差>1
		return false
	}
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
