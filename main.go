package main

import (
	"fmt"
	"strconv"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	a := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(a))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaqlgj/
// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))
	for _, str := range tokens {
		switch str {
		case "+":
			{
				num1 := stack.Pop()
				num2 := stack.Pop()
				stack.Push(num1 + num2)
			}
		case "-":
			{
				num1 := stack.Pop()
				num2 := stack.Pop()
				stack.Push(num2 - num1)
			}
		case "*":
			{
				num1 := stack.Pop()
				num2 := stack.Pop()
				stack.Push(num1 * num2)
			}
		case "/":
			{
				num1 := stack.Pop()
				num2 := stack.Pop()
				stack.Push(num2 / num1)
			}
		default:
			{
				num, _ := strconv.ParseInt(str, 10, 64)
				stack.Push(int(num))
			}
		}
	}
	return stack.Pop()
}

type Stack struct {
	cache  []int //存储数据，借助切片来排序
	length int   //数组大小
}

func NewStack(size int) *Stack {
	return &Stack{
		cache:  make([]int, 0, size),
		length: 0,
	}
}

// Push:往栈压入数据
func (s *Stack) Push(value int) {
	s.cache = append(s.cache, value)
	s.length++
}

// Pop:从栈取出数据
func (s *Stack) Pop() int {
	if s.length == 0 { //判断栈内是否有元素
		return 0
	}
	ret := s.cache[s.length-1]        //获取栈顶元素
	s.cache = s.cache[0 : s.length-1] //取出栈顶元素
	s.length--
	return ret
}
