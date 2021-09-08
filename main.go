package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 & 1)
	fmt.Println(2 & 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2, p3 := head, head.Next, head.Next //p1奇数队列的结尾,p2偶数队列的结尾,p3偶数队列的开头
	var idx int
	temp := head.Next.Next
	for temp != nil {
		idx++
		if idx&1 == 1 { //奇数
			p1.Next = temp
			p1 = temp
		} else { //偶数
			p2.Next = temp
			p2 = temp
		}
		temp = temp.Next
	}
	p2.Next = nil //偶数队列的最后一个要设置为nil
	p1.Next = p3
	return head
}
