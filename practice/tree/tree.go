package tree

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
