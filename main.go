package main

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
// 二叉搜索树中第K小的元素
//二叉搜索树中，左节点<根节点<右节点，所以中序遍历就能找出第k个最小值
func kthSmallest(root *TreeNode, k int) int {
	return inOrder(root, &k)
}

//中序遍历
func inOrder(root *TreeNode, k *int) int {
	if root != nil {
		val := inOrder(root.Left, k)
		if val != -1 {
			return val
		}
		*k--
		if *k == 0 {
			return root.Val
		}
		val = inOrder(root.Right, k)
		if val != -1 {
			return val
		}
	}
	return -1
}
