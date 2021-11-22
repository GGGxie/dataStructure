package main

import (
	"fmt"
<<<<<<< HEAD
	"math"
)

func main() {
	// n := uint32(0b00000010100101000001111010011100)
	fmt.Println(reverse(-123))
}

func hammingWeight(num uint32) int {
	var ret int
	for i := 0; i < 32; i++ {
		ret += int(num & 1)
		num >>= 1
	}
	return ret
}

// https://leetcode-cn.com/problems/reverse-integer/
// 整数反转
func reverse(x int) int {
	ret := 0
	for x != 0 {
		temp := x % 10
		x /= 10
		ret *= 10
		ret += temp
	}
	if ret > int(math.MaxInt32) || ret < int(-math.MaxInt32) {
		return 0
=======
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
>>>>>>> 2ca3180b6ab559c3de9dc77de3eafe604b4d4f35
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
