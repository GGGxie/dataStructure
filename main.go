package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
// 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapp := make(map[*ListNode]struct{})
	idx1, idx2 := headA, headB
	for idx1 != nil { //把headA指向的链表存入map中
		mapp[idx1] = struct{}{}
		idx1 = idx1.Next
	}
	for idx2 != nil { //找出headB指向的链表中第一个存在于map中的节点，这就是第一个公共节点
		if _, ok := mapp[idx2]; ok {
			return idx2
		}
		idx2 = idx2.Next
	}
	return nil
}
