package practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
// 删除某个链表中给定的（非末尾）节点
// 思路:直接用下一个节点的值来覆盖当前节点,使得达到"删除"效果
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xa0a97/
// 给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
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
