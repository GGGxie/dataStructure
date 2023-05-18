package minStack

import (
	"github.com/GGGxie/dataStructure/stacks/arrayStack"

	"math"
)

type MinStack struct {
	minCache *arrayStack.Stack //辅助栈，存的值与MinStack相同高度时的最小值。（用切片也能解决）
	cache    []int
	size     int
}

func Constructor() MinStack {
	stack := arrayStack.NewStack(1)
	stack.Push(math.MaxInt64) //插入一个最大值在栈底。方便第一次插入数据时比较
	return MinStack{
		minCache: stack,
		cache:    make([]int, 0),
		size:     0,
	}
}

func (this *MinStack) Push(val int) {
	this.cache = append(this.cache, val)
	this.size++
	//辅助栈插入最小值
	mc, _ := this.minCache.Peek()
	min := min(mc.(int), val)
	this.minCache.Push(min)
}

func (this *MinStack) Pop() {
	this.cache = this.cache[0 : this.size-1]
	this.size--
	//辅助栈移除一个最小值
	this.minCache.Pop()
}

func (this *MinStack) Top() int {
	return this.cache[this.size-1]
}

func (this *MinStack) GetMin() int {
	val, _ := this.minCache.Peek()
	return val.(int)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
