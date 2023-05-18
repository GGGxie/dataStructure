package main

import "fmt"

func main() {
	r := Constructor()
	r.Insert(&User{
		Id:     1,
		Number: 20,
	})
	r.Insert(&User{
		Id:     2,
		Number: 10,
	})
	r.Insert(&User{
		Id:     3,
		Number: 100,
	})
	r.Insert(&User{
		Id:     4,
		Number: 20,
	})
	r.Update(&User{
		Id:     4,
		Number: 200,
	})
	r.Delete(3)
	for _, u := range r.GetRank() {
		fmt.Println(u)
	}
}

//利用大堆实现动态实时更新的排行榜
func Constructor() *Rank {
	return &Rank{
		mh: &MaxHeap{
			Element: make([]*User, 0),
		},
	}
}

type Rank struct {
	mh *MaxHeap
}

//插入新数据
func (r *Rank) Insert(u *User) {
	r.mh.Push(u)
}

//删除数据
//时间复杂度O(n+logn)=O(n)
func (r *Rank) Delete(id int) {
	for i, u := range r.mh.Element {
		if u.Id == id {
			r.mh.Element = append(r.mh.Element[0:i], r.mh.Element[i+1:]...)
		}
	}
	//删掉之后重新生成大堆
	r.mh.Init()
}

//修改数据
//时间复杂度O(n+logn)=O(n)
func (r *Rank) Update(u *User) {
	for _, u2 := range r.mh.Element {
		if u2.Id == u.Id {
			u2.Number = u.Number
		}
	}
	//删掉之后重新生成大堆
	r.mh.Init()
}

//修改数据
//时间复杂度O(n+nlogn)=O(nlogn)
func (r *Rank) GetRank() []*User {
	length := len(r.mh.Element)
	ret := make([]*User, length)
	r.mh.Sort()
	for i, u := range r.mh.Element {
		ret[length-i-1] = &User{
			Id:     u.Id,
			Number: u.Number,
		}
	}
	return ret
}

type User struct {
	Id     int
	Number int
}

//大堆
type MaxHeap struct {
	Element []*User
}

// MaxHeap构造方法
func NewMaxHeap() *MaxHeap {
	// 第一个元素仅用于结束insert中的 for 循环
	h := &MaxHeap{Element: make([]*User, 0)}
	return h
}

//生成一个大堆
//时间复杂度O(logn)
func (H *MaxHeap) Init() {
	n := len(H.Element) - 1
	for i := n / 2; i >= 0; i-- {
		for { //下沉
			j1 := 2*i + 1
			if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
				break
			}
			//从两个子节点中选出一个大的
			j := j1 // left child
			if j2 := j1 + 1; j2 < n && H.Element[j2].Number > H.Element[j1].Number {
				j = j2 // = 2*i + 2  // right child
			}
			if H.Element[j].Number < H.Element[i].Number {
				break
			}
			H.Swap(i, j)
			i = j
		}
	}
}

// 插入元素,插入元素需要保证堆的性质
// 时间复杂度O(logn)
func (H *MaxHeap) Push(u *User) {
	H.Element = append(H.Element, u)
	j := len(H.Element) - 1
	for { //上浮插入的元素
		i := (j - 1) / 2 // parent
		if i == j || H.Element[i].Number > H.Element[j].Number {
			break
		}
		H.Swap(i, j)
		j = i
	}
}

// 删除并返回最大值
// 时间复杂度O(logn)
func (H *MaxHeap) Pop() (*User, error) {
	if len(H.Element) <= 1 {
		return nil, fmt.Errorf("MaxHeap is empty")
	}
	//取出切片首位元素
	maxElement := H.Element[0]
	//把最后一个元素挪到切片首位
	H.Swap(0, len(H.Element)-1)
	i, n := 0, len(H.Element)-1
	for { //下沉首位元素
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		//从两个子节点中选出一个大的
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && H.Element[j2].Number > H.Element[j1].Number {
			j = j2 // = 2*i + 2  // right child
		}
		if H.Element[j].Number < H.Element[i].Number {
			break
		}
		H.Swap(i, j)
		i = j
	}
	H.Element = H.Element[:n]
	return maxElement, nil
}

// 堆排序，对H内的切片进行排序
// 时间复杂度O(nlogn)
// 空间复杂度O(1)
func (H *MaxHeap) Sort() {
	n := len(H.Element) - 1
	for ; n > 0; n-- {
		i := 0
		H.Swap(0, n)
		for { //下沉
			j1 := 2*i + 1
			if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
				break
			}
			//从两个子节点中选出一个大的
			j := j1 // left child
			if j2 := j1 + 1; j2 < n && H.Element[j2].Number > H.Element[j1].Number {
				j = j2 // = 2*i + 2  // right child
			}
			if H.Element[j].Number < H.Element[i].Number {
				break
			}
			H.Swap(i, j)
			i = j
		}
	}
}

func (H *MaxHeap) Swap(i, j int) {
	H.Element[i], H.Element[j] = H.Element[j], H.Element[i]
}
