package heap

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/find-median-from-data-stream/
// 大小堆维护中位数
//首先想到用二叉搜索树来维护中位数,左右子树的节点个数相差<=1,根节点(或根节点+下一节点/2)便是中位数.(根节点放在右子树中)

type MedianFinder struct {
	MaxHp *minHp //左子树,用最小堆通过加入num的负数,取出时再取负来实现最大堆
	MinHp *minHp //右子树,
}

func Constructor() MedianFinder {
	return MedianFinder{
		MaxHp: &minHp{},
		MinHp: &minHp{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.MinHp.Len() == 0 || num > this.MinHp.Top().(int) { //入右子树
		heap.Push(this.MinHp, num)
		//确保两边子树的节点数量相差<=1
		if this.MinHp.Len() > this.MaxHp.Len()+1 {
			temp := heap.Pop(this.MinHp).(int)
			heap.Push(this.MaxHp, -temp)
		}
	} else {
		heap.Push(this.MaxHp, -num)
		//确保两边子树的节点数量相差<=1
		if this.MaxHp.Len() > this.MinHp.Len() {
			temp := heap.Pop(this.MaxHp).(int)
			heap.Push(this.MinHp, -temp)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	if this.MinHp.Len() > this.MaxHp.Len() { //节点为奇数
		return float64((this.MinHp.Top()).(int))
	} else { //节点为偶数
		return (float64(this.MinHp.Top().(int)) - float64(this.MaxHp.Top().(int))) / 2
	}
}

type minHp struct{ sort.IntSlice } //最小堆

func (h *minHp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *minHp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *minHp) Top() interface{} {
	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return nil
}
