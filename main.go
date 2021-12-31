package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个排序的链表
// https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	idx1, idx2 := l1, l2
	ret := &ListNode{} //ret指向新的链表的头结点，头结点不存数据，只为优化代码
	idx3 := ret
	for idx1 != nil && idx2 != nil {
		if idx1.Val < idx2.Val {
			idx3.Next = idx1
			idx1 = idx1.Next
		} else {
			idx3.Next = idx2
			idx2 = idx2.Next
		}
		idx3 = idx3.Next
	}
	if idx1 != nil {
		idx3.Next = idx1
	}
	if idx2 != nil {
		idx3.Next = idx2
	}
	return ret.Next //返回头结点.Next
}
