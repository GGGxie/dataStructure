package main

import "fmt"

func main() {
	A := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 1,
		},
		Right: nil,
	}
	T := A.Left
	fmt.Println(A, T)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof/submissions/
// 对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Judge(root.Left, root.Right)
}

//自底向上递归，根据对称性，left.Left=right.Right，left.Right=right.Left
func Judge(left, right *TreeNode) bool {
	if left == nil && right == nil { //两边为空，遍历结束
		return true
	} else if (left == nil || right == nil) || left.Val != right.Val { //单边为空，或者值不相等
		return false
	}
	return Judge(left.Left, right.Right) && Judge(left.Right, right.Left)
}
