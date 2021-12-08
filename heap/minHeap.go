package heap

import (
	"fmt"
	"math"
)

//heap中切片第一个元素为最大或最小值
type MinHeap struct {
	Element []int
}

// MinHeap构造方法
func NewMinHeap() *MinHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MinHeap{Element: []int{math.MinInt64}}
	return h
}

// 插入元素,插入元素需要保证堆的性质
func (H *MinHeap) Push(v int) {
	H.Element = append(H.Element, v)
	j := len(H.Element) - 1
	for {
		i := (j - 1) / 2 // parent
		if i == j || H.Element[i] < H.Element[j] {
			break
		}
		H.Element[i], H.Element[j] = H.Element[j], H.Element[i]
		j = i
	}
}

// 删除并返回最小值
// TODO还没验证
func (H *MinHeap) Pop() (int, error) {
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
