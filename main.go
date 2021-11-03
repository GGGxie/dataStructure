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

// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/submissions/
// 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return preOrder(root, p, q)
}

//先序遍历
func preOrder(root, p, q *TreeNode) *TreeNode {
	if root != nil {
		if root.Val == p.Val || root.Val == q.Val { //如果节点是左节点或者右节点则该返回
			return root
		}
		left := preOrder(root.Left, p, q)
		right := preOrder(root.Right, p, q)
		if left != nil && right != nil { //如果左和右都不为空，说明该节点就是适合返回的值
			return root
		}
		if left != nil { //如果左不为空，说明适合的值为左
			return left
		}
		if right != nil { //如果左不为空，说明适合的值为右
			return right
		}
	}
	return nil
}
