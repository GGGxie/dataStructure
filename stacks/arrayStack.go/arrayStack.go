package arrayStack

type Stack struct {
	cache []interface{}
	size  int
}

func (s *Stack) Push(value interface{}) {
	s.cache = append(s.cache, value)
	s.size++
}

func (s *Stack) Pop() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}
	ret := s.cache[0]
	if s.size == 1 {
		s.cache = nil
	} else {
		s.cache = s.cache[1:s.size]
	}
	s.size--
	return ret, true
}

func (s *Stack) Peek() (value interface{}, ok bool) {
	if s.size == 0 {
		return nil, false
	}
	return s.cache[0], true
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Clear() {
	s.cache = nil
	s.size = 0
}

//返回栈的值，通过copy的方式，不会对原值造成影响
func (s *Stack) Values() []interface{} {
	ret := make([]interface{}, s.size, s.size)
	copy(ret, s.cache)
	return ret
}
