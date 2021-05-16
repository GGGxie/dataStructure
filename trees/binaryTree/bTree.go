package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if node := lowestCommonAncestor(root.Left, p, q); node != nil {
		return node
	}
	if root == p || root == q {
		return root
	}
	if node := lowestCommonAncestor(root.Right, p, q); node != nil {
		return node
	}
	return nil
}
