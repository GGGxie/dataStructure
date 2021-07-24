package heap

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}

// MinHeap构造方法
func NewMinHeap() *MinHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

// 插入数字,插入数字需要保证堆的性质
func (H *MinHeap) Insert(v int) {
	H.Element = append(H.Element, v)
	i := len(H.Element) - 1
	// 上浮
	for ; H.Element[i/2] > v; i /= 2 {
		H.Element[i] = H.Element[i/2]
	}

	H.Element[i] = v
}

// 删除并返回最小值
func (H *MinHeap) DeleteMin() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	minElement := H.Element[1]
	lastElement := H.Element[len(H.Element)-1]
	var i, child int
	for i = 1; i*2 < len(H.Element); i = child {
		child = i * 2
		if child < len(H.Element)-1 && H.Element[child+1] < H.Element[child] {
			child++
		}
		// 下滤一层
		if lastElement > H.Element[child] {
			H.Element[i] = H.Element[child]
		} else {
			break
		}
	}
	H.Element[i] = lastElement
	H.Element = H.Element[:len(H.Element)-1]
	return minElement, nil
}
