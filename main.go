package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// 从上到下打印二叉树 II
type TempNode struct {
	node  *TreeNode
	Level int //记录层级
}
