package lists

type LinkList struct {
	Front *Node
	End   *Node
}

type Node struct {
	Data  int
	Next  *Node
	Front *Node
}

func (l *LinkList) InitList() *LinkList {
	return &LinkList{}
}

//insert 插入新节点
func (l *LinkList) Insert() {

}
