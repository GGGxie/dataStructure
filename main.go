package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TempNode struct {
	node  *TreeNode
	Level int //记录层级
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var list []*TempNode
	list = append(list, &TempNode{
		node:  root,
		Level: 1,
	})
	for len(list) != 0 { //不能用list!=nil判断，empty slice != nil
		top := list[0]            //取出头部元素
		list = list[1:]           //弹出头部元素
		if top.node.Left != nil { //加入左节点
			list = append(list, &TempNode{
				node:  top.node.Left,
				Level: top.Level + 1,
			})
		}
		if top.node.Right != nil { //加入右节点
			list = append(list, &TempNode{
				node:  top.node.Right,
				Level: top.Level + 1,
			})
		}
		if len(ret) < top.Level { //扩容
			ret = append(ret, []int{})
		}
		ret[top.Level-1] = append(ret[top.Level-1], top.node.Val)
	}
	return ret
}
