package stack

import (
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/min-stack/
// 最小栈
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type MinStack struct {
	minCache []int //辅助队列，入栈一个新值，便加一个对应的最小值
	cache    []int
	size     int
}

func Constructor() MinStack {
	mc := make([]int, 1)
	mc[0] = math.MaxInt64
	return MinStack{
		minCache: mc,
		cache:    make([]int, 0),
		size:     0,
	}
}

func (this *MinStack) Push(val int) {
	this.cache = append(this.cache, val)
	this.size++
	//辅助栈插入最小值
	min := min(this.minCache[len(this.minCache)-1], val)
	this.minCache = append(this.minCache, min)
}

func (this *MinStack) Pop() {
	this.cache = this.cache[0 : this.size-1]
	this.size--
	//辅助栈移除一个最小值
	this.minCache = this.minCache[0 : len(this.minCache)-1]
}

func (this *MinStack) Top() int {
	return this.cache[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.minCache[len(this.minCache)-1]
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaqlgj/
// 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := NewStack(len(tokens))
	for _, str := range tokens {
		switch str { //遇到操作符就取出栈中两个数字进行计算
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
		default: //遇到数字就压入栈中
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
