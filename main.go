package main

func main() {
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mapp := make(map[*Node]*Node)
	var A *Node
	temp := head
	for temp != nil {
		node := &Node{
			Val: head.Val,
		}
		mapp[temp] = head
		temp = temp.Next
	}

	return temp
}
