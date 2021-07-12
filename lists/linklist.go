package lists

import "fmt"

type LinkList struct {
	Header *Node
	End    *Node
	Length int //链表空间大小
	Size   int //链表现有数据量
}

type Node struct {
	Data  int
	Front *Node
	Next  *Node
}

func InitList() *LinkList {
	header := &Node{
		Data:  0,
		Front: nil,
		Next:  nil,
	}
	return &LinkList{
		Header: header,
		End:    header,
	}
}

//insert 在链头插入新元素
func (l *LinkList) Add(data int) {
	temp := &Node{
		Data:  data,
		Front: nil,
		Next:  l.Header.Next,
	}
	l.Header.Next.Front = temp
	l.Header.Next = temp
	l.Size++
}

//insert 在链尾插入新元素
func (l *LinkList) Append(data int) {
	if l.Size == 0 {
		temp := &Node{
			Data:  data,
			Front: nil,
			Next:  nil,
		}
		l.Header = temp
		l.End = temp
	} else {
		temp := &Node{
			Data:  data,
			Front: l.End,
			Next:  nil,
		}
		l.End.Next = temp
		l.End = temp
	}
	l.Size++
}

//在链表第i个元素后插入新元素
func (l *LinkList) Insert(data int, i int) {
	if i >= l.Length { //超过链表长度,就在尾部插入
		l.Append(data)
		return
	}
	tempNode := l.Get(i)
	newNode := &Node{
		Data:  data,
		Front: tempNode,
		Next:  tempNode.Next,
	}
	tempNode.Next = newNode
}

// 获取链表中下标为index的元素
func (l *LinkList) Get(index int) *Node {
	if index > l.Length {
		return nil
	}
	temp := l.Header
	for i := 1; i < index; i++ {
		temp = temp.Next
	}
	return temp
}

func (l *LinkList) Iterate() {
	temp := l.Header.Next
	for temp != nil {
		fmt.Println(temp)
		temp = temp.Next
	}
	fmt.Println(l.End)
}
