package binarySearchTree

import "math"

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

//给一个有序数组，生成一个二分搜索树
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

type ListNode struct {
	Val  int
	Next *ListNode
}

//bfs 层序遍历，广度搜索
func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}
	var tempList []*TreeNode //存储数据的队列
	var retList []*ListNode  //返回的数据
	tempList = append(tempList, tree)
	for len(tempList) > 0 {
		var index, tempNode *ListNode
		size := len(tempList)
		for i := 0; i < size; i++ {
			if tempList[i].Left != nil {
				tempList = append(tempList, tempList[i].Left)
			}
			if tempList[i].Right != nil {
				tempList = append(tempList, tempList[i].Right)
			}
			if i == 0 {
				tempNode = &ListNode{
					Val:  tempList[i].Val,
					Next: nil,
				}
				index = tempNode
			} else {
				index.Next = &ListNode{
					Val:  tempList[i].Val,
					Next: nil,
				}
				index = index.Next
			}
		}
		retList = append(retList, tempNode)
		tempList = tempList[size:]
		size = len(tempList)
	}
	return retList
}

//层序遍历，判断是否为平衡二叉树，记录深度和节点个数，判断节点个数是否>2^(depth-1)-1
func isBalanced(tree *TreeNode) bool {
	return Height(tree) >= 0
}

//返回节点的高度,如果以该节点为根节点的二叉树,不为平衡二叉树,返回-1
func Height(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	leftHight := Height(tree.Left)
	if leftHight == -1 { //优化点:左节点不平衡就不需要继续优化
		return -1
	}
	rightHeight := Height(tree.Right)
	if rightHeight == -1 || abs(leftHight-rightHeight) > 1 {
		return -1
	}
	return max(leftHight, rightHeight) + 1
}

//判断是否二叉搜索树
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

//中序找出后继节点,要点:这是有序的,比p大的就是它的后续节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node := inorderSuccessor(root.Left, p); node != nil {
		return node
	}
	if root.Val > p.Val {
		return root
	}
	if node := inorderSuccessor(root.Right, p); node != nil {
		return node
	}
	return nil
}
