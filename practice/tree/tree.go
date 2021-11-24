package tree

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/flatten-nested-list-iterator/
// 扁平化嵌套列表迭代器
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

// type NestedIterator struct {
// 	vals []int
// }

// // nestedList看成一棵多根节点树，NestedInteger要么是叶子节点，要么是非叶子节点，非叶子用dfs遍历，叶子节点值加入迭代器数组
// func Constructor(nestedList []*NestedInteger) *NestedIterator {
// 	var vals []int
// 	var dfs func(ni []*NestedInteger)
// 	dfs = func(ni []*NestedInteger) {
// 		for _, nestedInteger := range ni {
// 			if nestedInteger.IsInteger() {
// 				vals = append(vals, nestedInteger.GetInteger())
// 			} else {
// 				dfs(nestedInteger.GetList())
// 			}
// 		}
// 	}
// 	dfs(nestedList)
// 	return &NestedIterator{ //把得到的所有叶子节点值放进迭代器数组
// 		vals: vals,
// 	}
// }

// //每次返回第一个，并把第一个值排出对垒
// func (it *NestedIterator) Next() int {
// 	val := it.vals[0]
// 	it.vals = it.vals[1:]
// 	return val
// }

// func (it *NestedIterator) HasNext() bool {
// 	return len(it.vals) > 0
// }

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

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// 二叉树的序列化与反序列化
type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var ret []string
	var preOrder func(root *TreeNode, str *[]string)
	preOrder = func(root *TreeNode, str *[]string) {
		if root != nil {
			*str = append(*str, strconv.FormatInt(int64(root.Val), 10))
			preOrder(root.Left, str)
			preOrder(root.Right, str)
			return
		}
		*str = append(*str, "null")
	}
	preOrder(root, &ret)
	return strings.Join(ret, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	tl := strings.Split(data, ",")
	var build func(tl *[]string) *TreeNode
	build = func(tl *[]string) *TreeNode { //前序构建，先构建根节点，再递归构建左节点，递归构建右节点
		tn := (*tl)[0]
		*tl = (*tl)[1:]
		if tn == "null" {
			return nil
		} else {
			val, _ := strconv.ParseInt(tn, 10, 64)
			return &TreeNode{ //序列化的顺序是前序遍历，反序列化也需要相同顺序，中、左、右
				Val:   int(val),
				Left:  build(tl),
				Right: build(tl),
			}
		}
	}
	return build(&tl)
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

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
// 从上到下打印二叉树 II
type TempNode struct {
	node  *TreeNode
	Level int //记录层级
}

func levelOrder2(root *TreeNode) [][]int {
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

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
// 从上到下打印二叉树 III
func levelOrder3(root *TreeNode) [][]int {
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
		top := list[0]  //取出头部元素
		list = list[1:] //弹出头部元素
		if top.node == nil {
			continue
		}
		list = append(list, &TempNode{ //加入左节点
			node:  top.node.Left,
			Level: top.Level + 1,
		})
		list = append(list, &TempNode{ //加入右节点
			node:  top.node.Right,
			Level: top.Level + 1,
		})

		if len(ret) < top.Level { //扩容
			ret = append(ret, []int{})
		}
		if top.Level&1 == 1 { //奇数行，从左往右遍历
			ret[top.Level-1] = append(ret[top.Level-1], top.node.Val)
		} else { //偶数行，从右往左遍历
			ret[top.Level-1] = append([]int{top.node.Val}, ret[top.Level-1]...)
		}
	}
	return ret
}

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/
// 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil { //B为空返回false
		return false
	}
	return PreOrder(A, B)
}

//前序遍历，每个节点单独和B进行匹配判断B是否为A的子树
func PreOrder(root *TreeNode, B *TreeNode) bool {
	if ok := Compare(root, B); ok { //用root和B进行比较
		return true
	}
	if root != nil {
		if ok := PreOrder(root.Left, B); ok { //用root.Left和B进行比较
			return true
		}
		if ok := PreOrder(root.Right, B); ok { //用root.Right和B进行比较
			return true
		}
	}
	return false
}

//节点匹配
func Compare(root *TreeNode, B *TreeNode) bool { //递归比较从root节点开始和B是不是同一个子树
	if B == nil { //B为空则递归结束，B为A的子树
		return true
	} else if root == nil && B != nil { //A已经没有节点了，而B还没遍历完成，B不是A的子树
		return false
	}
	if root.Val == B.Val { //判断两个节点是否相等
		if ok := Compare(root.Left, B.Left); !ok { //继续递归比较左节点
			return false
		}
		if ok := Compare(root.Right, B.Right); !ok { //继续递归比较右节点
			return false
		}
	} else { //不相等，B不是A的子树
		return false
	}
	return true
}
