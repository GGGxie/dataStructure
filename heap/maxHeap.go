package heap

import (
	"container/heap"
	"fmt"
)

//实现heap包
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *MaxHeap) Top() interface{} {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return nil
}

func Test() {
	h := make(MaxHeap, 0)
	heap.Init(&h)
	h = []int{62, 41, 30, 28, 16, 22, 13, 19, 17, 15}
	heap.Push(&h, 52)
	fmt.Println(h)
	// heap.Push(&h, 1)
	// heap.Push(&h, 4)
	// heap.Push(&h, 5)
	// heap.Push(&h, 2)

	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
}
