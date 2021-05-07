package binarySearchTree

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
