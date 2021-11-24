package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
