package main

import (
	"fmt"
)

func main() {
	z := []int{0, 1, 2, 3}
	fmt.Println(minArray(z))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
// 从上到下打印二叉树(层序遍历)
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	var list []*TreeNode
	list = append(list, root)
	for len(list) != 0 { //不能用list!=nil判断，empty slice != nil
		top := list[0]       //取出头部元素
		list = list[1:]      //弹出头部元素
		if top.Left != nil { //加入左节点
			list = append(list, top.Left)
		}
		if top.Right != nil { //加入右节点
			list = append(list, top.Right)
		}
		ret = append(ret, top.Val)
	}
	return ret
}

//先序遍历
func PreOrder(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	*ret = append(*ret, root.Val)
	PreOrder(root.Left, ret)
	PreOrder(root.Right, ret)
}
