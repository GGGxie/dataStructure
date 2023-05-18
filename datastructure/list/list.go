package list

import "fmt"

// 链表结构
type LinkList struct {
	Header *Node
	End    *Node
	Length int //链表空间大小
	Size   int //链表现有数据量
}

// 双向节点结构
type Node struct {
	Data  int
	Front *Node
	Next  *Node
}

func NewNode(data int, front, next *Node) *Node {
	return &Node{
		Data:  data,
		Front: front,
		Next:  next,
	}
}

// init 初始化一个新链表
func InitList() *LinkList {
	return &LinkList{
		Header: nil,
		End:    nil,
		Length: 0,
		Size:   0,
	}
}

// insert 在链头插入新元素
func (l *LinkList) Add(data int) {
	if l.Size == 0 { //空链表
		temp := NewNode(data, nil, nil)
		l.Header = temp
		l.End = temp
	} else {
		temp := NewNode(data, nil, l.Header)
		l.Header.Front = temp
		l.Header = temp
	}
	l.Size++
}

// insert 在链尾插入新元素
func (l *LinkList) Append(data int) {
	if l.Size == 0 { //空链表
		temp := NewNode(data, nil, nil)
		l.Header = temp
		l.End = temp
	} else {
		temp := NewNode(data, l.End, nil)
		l.End.Next = temp
		l.End = temp
	}
	l.Size++
}

// 在链表第i个元素后插入新元素
func (l *LinkList) Insert(data int, i int) {
	if i > l.Size { //超过链表长度,就在尾部插入
		l.Append(data)
		return
	}
	tempNode := l.Get(i) //获取第i个元素
	newNode := NewNode(data, tempNode, tempNode.Next)
	tempNode.Next.Front = newNode
	tempNode.Next = newNode
	l.Size++
}

// 获取链表第index个元素
func (l *LinkList) Get(index int) *Node {
	if index > l.Size { //超过链表长度
		return nil
	}
	temp := l.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}
	return temp
}

// 删除第index个元素，返回被删除元素
func (l *LinkList) Delete(index int) *Node {
	if index > l.Size { //超过链表长度
		return nil
	}
	var temp *Node
	if index == 1 { //删除第一个元素
		temp = l.Header
		l.Header.Next.Front = nil
		l.Header = l.Header.Next
	} else if index == l.Size { //删除最后一个元素
		temp = l.End
		l.End.Front.Next = nil
		l.End = l.End.Front
	} else { //删除中间元素
		temp = l.Get(index)
		temp.Front.Next = temp.Next
		temp.Next.Front = temp.Front
	}
	l.Size--
	return temp
}

// 正序打印链表
func (l *LinkList) Iterate() {
	if l.Size == 0 {
		return
	}
	temp := l.Header
	for temp != nil {
		fmt.Println(temp)
		temp = temp.Next
	}
}

// 倒序打印链表
func (l *LinkList) Reverse() {
	if l.Size == 0 {
		return
	}
	temp := l.End
	for temp != nil {
		fmt.Println(temp)
		temp = temp.Front
	}
}
