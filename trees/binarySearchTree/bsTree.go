package binarySearchTree

import "fmt"

type Tree struct {
	root *Node
	size int
}

type Node struct {
	value int
	left  *Node
	right *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给一个数组，生成一个二分搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	len := len(nums)
	if len == 0 {
		return nil
	}
	mid := len / 2
	tree := &TreeNode{
		Val:   nums[mid],
		Left:  dfs(nums, 0, mid-1),
		Right: dfs(nums, mid+1, len-1),
	}
	return tree
}

//二分递归
func dfs(nums []int, start, end int) *TreeNode {
	if end < start {
		return nil
	}
	mid := start + (end-start)/2
	if start == end {
		return &TreeNode{
			Val:   nums[mid],
			Left:  nil,
			Right: nil,
		}
	} else {
		return &TreeNode{
			Val:   nums[mid],
			Left:  dfs(nums, start, mid-1),
			Right: dfs(nums, mid+1, end),
		}
	}
}

type Stack struct {
	cache []interface{}
	size  int
}

func (s *Stack) Push(value interface{}) {
	if value == nil {
		return
	}
	s.cache = append(s.cache, value)
	s.size++
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}
	ret := s.cache[0]
	if s.size == 1 {
		s.cache = nil
	} else {
		s.cache = s.cache[1:s.size]
	}
	s.size--
	return ret, true
}

func (s *Stack) Peek() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}
	return s.cache[0], true
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	var stack Stack
	var tempList []*TreeNode
	if tree == nil {
		return nil
	}
	tempList = append(tempList, tree)
	if tree.Left != nil {
		stack.Push(tree.Left)
	}
	if tree.Right != nil {
		stack.Push(tree.Right)
	}
	for !stack.Empty() { //获取数组
		if tree, ok := stack.Pop(); ok {
			tempTree := tree.(*TreeNode)
			tempList = append(tempList, tempTree)
			if tempTree.Left != nil {
				stack.Push(tempTree.Left)
			}
			if tempTree.Right != nil {
				stack.Push(tempTree.Right)
			}
		}
	}
	return getListOfDepth(tempList)
}

func getListOfDepth(temp []*TreeNode) (ret []*ListNode) {
	var tempListNode *ListNode
	var endListNode *ListNode
	var col int
	for i := range temp {
		fmt.Println(col, i, temp[i])
		if temp[i] == nil {
			continue
		}
		if i == 0 || i > ((2<<(col-1))-1) { //起始点||到达下一层
			if temp[i] == nil {
				tempListNode = &ListNode{
					Val:  temp[i].Val,
					Next: nil,
				}
			}
			tempListNode = &ListNode{
				Val:  temp[i].Val,
				Next: nil,
			}
			ret = append(ret, tempListNode) //加入数组
			endListNode = tempListNode
			col++
		} else {
			endListNode.Next = &ListNode{
				Val:  temp[i].Val,
				Next: nil,
			}
			endListNode = endListNode.Next
		}
	}
	return ret
}
