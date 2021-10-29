package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	fmt.Println(maxSlidingWindow([]int{7, 4, 2, 3, 1}, 2))

}

var tempNums []int

type maxHp struct{ sort.IntSlice } //最大堆，存的是数组的下表
func (h *maxHp) Less(i, j int) bool { //根据数组的值由大到小排序
	return tempNums[h.IntSlice[i]] > tempNums[h.IntSlice[j]]
}

func (h *maxHp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *maxHp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *maxHp) Top() interface{} {
	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return nil
}

func maxSlidingWindow(nums []int, k int) []int {
	tempNums = nums
	maxHp := &maxHp{}
	for i := 0; i < k; i++ {
		heap.Push(maxHp, i)
	}
	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[maxHp.Top().(int)]
	for i := k; i < n; i++ {
		heap.Push(maxHp, i)
		for maxHp.Top().(int) <= i-k { //把在栈顶，且不在窗口内的元素全部移除
			heap.Pop(maxHp)
		}
		ans = append(ans, nums[maxHp.Top().(int)])
	}
	return ans
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	maxHp := &maxHp{}
// 	var ret []int
// 	var i int
// 	for i = 0; i < k; i++ {
// 		heap.Push(maxHp, nums[i])
// 	}
// 	ret = append(ret, maxHp.Top().(int))
// 	for j := i; j < len(nums); j++ {
// 		//移除堆顶元素
// 		maxHp.Pop()
// 		//添加新元素
// 		heap.Push(maxHp, nums[j])
// 		//加入ret
// 		ret = append(ret, maxHp.Top().(int))
// 	}
// 	return ret
// }

// 对于「最大值」，我们可以想到一种非常合适的数据结构，那就是优先队列（堆），其中的大根堆可以帮助我们实时维护一系列元素中的最大值。

// 对于本题而言，初始时，我们将数组 \textit{nums}nums 的前 kk 个元素放入优先队列中。每当我们向右移动窗口时，我们就可以把一个新的元素放入优先队列中，此时堆顶的元素就是堆中所有元素的最大值。然而这个最大值可能并不在滑动窗口中，在这种情况下，这个值在数组 \textit{nums}nums 中的位置出现在滑动窗口左边界的左侧。因此，当我们后续继续向右移动窗口时，这个值就永远不可能出现在滑动窗口中了，我们可以将其永久地从优先队列中移除。

// 我们不断地移除堆顶的元素，直到其确实出现在滑动窗口中。此时，堆顶元素就是滑动窗口中的最大值。为了方便判断堆顶元素与滑动窗口的位置关系，我们可以在优先队列中存储二元组 (\textit{num}, \textit{index})(num,index)，表示元素 \textit{num}num 在数组中的下标为 \textit{index}index。
