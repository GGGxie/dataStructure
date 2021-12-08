package main

import (
	"fmt"
	"math"
)

func main() {
	h := MaxHeap{
		Element: []int{62, 41, 30, 28, 16, 22, 13, 19, 17, 15},
	}
	// heap.Init(&h)
	h.Push(52)
	fmt.Println(h)
	h.Pop()
	fmt.Println(h)
	h.Sort()
	fmt.Println(h)
	// heap.Push(&h, 4)
	// heap.Push(&h, 5)
	// heap.Push(&h, 2)

	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))
	// fmt.Println(heap.Pop(&h))

}

//大堆
type MaxHeap struct {
	Element []int
}

// MaxHeap构造方法
func NewMaxHeap() *MaxHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MaxHeap{Element: []int{math.MinInt64}}
	return h
}

// 插入元素,插入元素需要保证堆的性质
// 时间复杂度O(logn)
func (H *MaxHeap) Push(v int) {
	H.Element = append(H.Element, v)
	j := len(H.Element) - 1
	for { //上浮插入的元素
		i := (j - 1) / 2 // parent
		if i == j || H.Element[i] > H.Element[j] {
			break
		}
		H.Swap(i, j)
		j = i
	}
}

// 删除并返回最大值
// 时间复杂度O(logn)
func (H *MaxHeap) Pop() (int, error) {
	if len(H.Element) <= 1 {
		return 0, fmt.Errorf("MaxHeap is empty")
	}
	//取出切片首位元素
	maxElement := H.Element[0]
	//把最后一个元素挪到切片首位
	H.Swap(0, len(H.Element)-1)
	i, n := 0, len(H.Element)-1
	for { //下沉首位元素
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		//从两个子节点中选出一个大的
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && H.Element[j2] > H.Element[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		if H.Element[j] < H.Element[i] {
			break
		}
		H.Swap(i, j)
		i = j
	}
	H.Element = H.Element[:n]
	return maxElement, nil
}

// 堆排序，对H内的切片进行排序
// 时间复杂度O(nlogn)
func (H *MaxHeap) Sort() {
	n := len(H.Element) - 1
	for i := n/2 - 1; i >= 0; i-- {
		for { //下沉
			j1 := 2*i + 1
			if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
				break
			}
			//从两个子节点中选出一个大的
			j := j1 // left child
			if j2 := j1 + 1; j2 < n && H.Element[j2] > H.Element[j1] {
				j = j2 // = 2*i + 2  // right child
			}
			if H.Element[j] < H.Element[i] {
				break
			}
			H.Swap(i, j)
			i = j
		}
	}
}

func (H *MaxHeap) Swap(i, j int) {
	H.Element[i], H.Element[j] = H.Element[j], H.Element[i]
}
