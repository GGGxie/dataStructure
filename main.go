package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseList(nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

type Stack struct {
	cache  []*ListNode //存储数据，借助切片来排序
	length int         //数组大小
}

func NewStack(size int) *Stack {
	return &Stack{
		cache:  make([]*ListNode, 0, size),
		length: 0,
	}
}

// Push:往栈压入数据
func (s *Stack) Push(value *ListNode) {
	if value == nil {
		return
	}
	s.cache = append(s.cache, value)
	s.length++
}

// Pop:从栈取出数据
func (s *Stack) Pop() *ListNode {
	if s.length == 0 { //判断栈内是否有元素
		return nil
	}
	ret := s.cache[s.length-1]        //获取栈顶元素
	s.cache = s.cache[0 : s.length-1] //取出栈顶元素
	s.length--
	return ret
}

// Empty:判断栈是否为空
func (s *Stack) Empty() bool {
	return s.length == 0
}
