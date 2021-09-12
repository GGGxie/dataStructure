package practice

import "sort"

// 链表相关练习

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xam1wr/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//链表深拷贝
func copyRandomList(head *Node) *Node {
	mapp := make(map[*Node]*Node) //记录新旧节点之间的关系
	var ret *Node
	temp := head
	for temp != nil {
		node := &Node{
			Val: temp.Val,
		}
		mapp[temp] = node
		temp = temp.Next
	}
	//再次遍历
	temp = head
	for temp != nil {
		mapp[temp].Next = mapp[temp.Next]
		mapp[temp].Random = mapp[temp.Random]
		temp = temp.Next
	}
	ret = mapp[head]
	return ret
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaazns/
// 给定一个链表，判断链表中是否有环。
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

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaazns/
// 判断链表是否有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil { //只需要判断快指针的next是否为空，不需要判断慢指针
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xa262d/
// 链表排序
type ll []*ListNode

func (a ll) Len() int           { return len(a) }
func (a ll) Less(i, j int) bool { return a[i].Val < a[j].Val }
func (a ll) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var slice ll
	temp := head
	for temp != nil { //记录元素和节点的对应关系
		slice = append(slice, temp)
		temp = temp.Next
	}
	//给map排序
	sort.Sort(slice)
	for i := range slice {
		if i != len(slice)-1 {
			slice[i].Next = slice[i+1]
		} else {
			slice[i].Next = nil
		}
	}
	head = slice[0]
	return head
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xan8ah/
// 找出两条单链表的第一个相交节点，没有则返回nil
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xavip3/
// 翻转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var ll []*ListNode //借助数组来记录顺序
	temp := head
	for temp != nil {
		ll = append(ll, temp)
		temp = temp.Next
	}
	for i := len(ll) - 1; i >= 0; i-- {
		if i == 0 {
			ll[i].Next = nil
		} else {
			ll[i].Next = ll[i-1]
		}
	}
	return ll[len(ll)-1] //返回翻转后的头部
}
