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
	if head == nil {
		return head
	}
	odd := head           //奇数列表尾部
	evenHead := head.Next //偶数列表头部
	even := evenHead      //偶数列表尾部
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = even.Next
		even.Next = odd.Next
		even = odd.Next
	}
	odd.Next = evenHead
	return head
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
}
