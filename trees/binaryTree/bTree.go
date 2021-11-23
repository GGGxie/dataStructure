package binaryTree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层序遍历，判断是否为平衡二叉树
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

//找出二叉树中某两个节点的第一个共同祖先
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

//判断t2是否为t1的子树
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1.Val == t2.Val { //注意,值相等就行,地址不一定相等,题目有问题?应该是地址相等才对啊!!!!!!
		return true
	}
	if t1.Left != nil {
		if flag := checkSubTree(t1.Left, t2); flag {
			return true
		}
	}
	if t1.Right != nil {
		if flag := checkSubTree(t1.Right, t2); flag {
			return true
		}
	}
	return false
}

var count int

//输出总和为sum的路径总数
func pathSum(root *TreeNode, sum int) int {
	rootDfs(root, sum)
	return count
}

func rootDfs(root *TreeNode, sum int) {
	setRoot(root, sum)
	if root.Left != nil {
		rootDfs(root.Left, sum)
	}
	if root.Right != nil {
		rootDfs(root.Right, sum)
	}
}

func setRoot(root *TreeNode, sum int) {
	dfs(root, 0, sum)
}

func dfs(root *TreeNode, pathSum, sum int) {
	if root.Left == nil && root.Right == nil {
		fmt.Println(root.Val)
		pathSum += root.Val
		if pathSum == sum {
			count++
		}
		return
	}
	pathSum += root.Val
	if root.Left != nil {
		dfs(root.Left, pathSum, sum)
	}
	if root.Right != nil {
		dfs(root.Right, pathSum, sum)
	}
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
