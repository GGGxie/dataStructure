package sortStack

import "github.com/GGGxie/dataStructure/datastructure/stacks/arrayStack"

// 排序栈
// 辅助数据结构：普通栈
// 最小的元素在栈顶
type SortedStack struct {
	stack *arrayStack.Stack
	cache []interface{}
	size  int
}

func Constructor() SortedStack {
	s := SortedStack{
		stack: &arrayStack.Stack{},
	}
	return s
}

func (s *SortedStack) Push(val interface{}) {
	if s.size != 0 { //把所有比val小的值先迁移到辅助stack
		for compare(s.Peek(), val) < 0 {
			if s.size == 0 {
				break
			}
			temp, ok := s.Pop()
			if !ok {
				return
			}
			s.stack.Push(temp)
		}
	}
	s.cache = append(s.cache, val)
	s.size++
	for !s.stack.Empty() { //辅助stack的数据都迁移回来
		temp, ok := s.stack.Pop()
		if !ok {
			return
		}
		s.cache = append(s.cache, temp)
		s.size++
	}
}

func (s *SortedStack) Pop() (value interface{}, ok bool) {
	if s.size == 0 {
		return
	}
	s.size--
	if s.size == 0 {
		s.cache = nil
	} else {
		s.cache = s.cache[0:s.size]
	}
	return
}

func (s *SortedStack) Peek() interface{} {
	if s.size == 0 {
		return nil
	}
	ret := s.cache[s.size-1]
	return ret
}

func (s *SortedStack) IsEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false

}
func (s *SortedStack) Empty() bool {
	return s.size == 0
}

func (s *SortedStack) Size() int {
	return s.size
}

func (s *SortedStack) Clear() {
	s.cache = nil
	s.size = 0
}

// 返回栈的值，通过copy的方式，不会对原值造成影响
func (s *SortedStack) Values() []interface{} {
	ret := make([]interface{}, s.size, s.size)
	copy(ret, s.cache)
	return ret
}

func compare(a, b interface{}) int {
	if a.(int) == b.(int) {
		return 0
	} else if a.(int) < b.(int) {
		return -1
	} else {
		return 1
	}
}
