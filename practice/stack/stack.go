package stack

import "math"

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
