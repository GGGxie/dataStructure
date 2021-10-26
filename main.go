package main

import (
	"container/heap"
	"sort"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
}
func kthSmallest(matrix [][]int, k int) int {
	minHeap := &hp{}
	for i := range matrix {
		heap.Push(minHeap, matrix[i])
	}
	for i := 0; i < k-1; i++ {
		heap.Pop(minHeap)
	}
	return heap.Pop(minHeap).(int)
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
