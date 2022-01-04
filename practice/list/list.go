package list

import (
	"container/list"
	"sort"
)

// 链表相关练习

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xam1wr/
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//链表深拷贝
func copyRandomList(head *Node) *Node {
	mapp := make(map[*Node]*Node) //记录新旧节点之间的关系。key：旧节点，value：新节点
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

// // https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
// // 两个链表的第一个公共节点
// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	mapp := make(map[*ListNode]struct{})
// 	idx1, idx2 := headA, headB
// 	for idx1 != nil { //把headA指向的链表存入map中
// 		mapp[idx1] = struct{}{}
// 		idx1 = idx1.Next
// 	}
// 	for idx2 != nil { //找出headB指向的链表中第一个存在于map中的节点，这就是第一个公共节点
// 		if _, ok := mapp[idx2]; ok {
// 			return idx2
// 		}
// 		idx2 = idx2.Next
// 	}
// 	return nil
// }

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

// // https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof/
// // 翻转链表
// // 空间复杂度O(1)
// func reverseList(head *ListNode) *ListNode {
// 	cur := head        //记录当前节点
// 	var prev *ListNode //记录前一个节点
// 	for cur != nil {
// 		next := cur.Next //记录后一个节点
// 		cur.Next = prev  //当前指向前缀
// 		prev = cur       //prev=当前
// 		cur = next       //当前=后一个节点
// 	}
// 	return prev
// }

//链表实现LRU页面置换算法
type LRUCache struct {
	capacity int                   //容量
	cache    map[int]*list.Element //存储
	list     *list.List            //利用链表来给元素排序
}

type Pair struct { //元素结构体定义
	key   int
	value int
}

//新建一个缓存
func Constructor(c int) LRUCache {
	return LRUCache{
		capacity: c,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

//根据key获取数据，并把数据放到链表表头
func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

//插入新数据，并把数据放到链表表头
func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		delete(this.cache, key)
		this.list.Remove(elem)
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	} else {
		if len(this.cache) >= this.capacity {
			delete(this.cache, this.list.Back().Value.(Pair).key)
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(Pair{key, value})
		this.cache[key] = this.list.Front()
	}
}

// https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof/
// 删除链表的节点
func deleteNode2(head *ListNode, val int) *ListNode {
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		if cur.Val == val {
			if pre == nil { //删除头结点
				head = head.Next
			} else { //删除中间节点
				pre.Next = cur.Next
			}
			break
		}
		pre = cur
		cur = cur.Next
	}
	return head
}

// https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/submissions/
// 链表中倒数第k个节点
//快慢指针，快指针-慢指针=k
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var idx1, idx2 *ListNode //idx1：慢指针，idx2：快指针
	idx1, idx2 = head, head
	for i := 0; i < k; i++ { //让idx2先走k步
		idx2 = idx2.Next
	}
	for idx2 != nil {
		idx1 = idx1.Next
		idx2 = idx2.Next
	}
	return idx1
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

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 调整数组顺序使奇数位于偶数前面
// 双指针法
func exchange(nums []int) []int {
	length := len(nums)
	idx1, idx2 := 0, length-1 //idx1指向奇数，idx2指向偶数
	for {
		for { //for循环找到奇数
			if idx1 >= length || nums[idx1]&1 == 0 {
				break
			}
			idx1++
		}
		for { //for循环找到偶数
			if idx2 < 0 || nums[idx2]&1 == 1 {
				break
			}
			idx2--
		}
		if idx1 >= length || idx2 < 0 || idx1 >= idx2 { //下标判断
			break
		}
		nums[idx1], nums[idx2] = nums[idx2], nums[idx1] //交换
	}
	return nums
}
