package arrayStack

type Stack struct {
	cache  []interface{} //存储数据，借助切片来排序
	length int           //数组大小
}

func NewStack(size int) *Stack {
	return &Stack{
		cache:  make([]interface{}, 0, size),
		length: 0,
	}
}

// Push:往栈压入数据
func (s *Stack) Push(value interface{}) {
	if value == nil {
		return
	}
	s.cache = append(s.cache, value)
	s.length++
}

// Pop:从栈取出数据
func (s *Stack) Pop() (value interface{}, ok bool) {
	if s.length == 0 { //判断栈内是否有元素
		return nil, false
	}
	ret := s.cache[s.length-1]        //获取栈顶元素
	s.cache = s.cache[0 : s.length-1] //取出栈顶元素
	s.length--
	return ret, true
}

// Peek:获取栈顶元素的值
func (s *Stack) Peek() (value interface{}, ok bool) {
	if s.length == 0 { //判断栈是否为空
		return nil, false
	}
	return s.cache[0], true
}

// Size:获取栈的大小
func (s *Stack) Size() int {
	return s.length
}

// Clear:清空栈
func (s *Stack) Clear() {
	s.cache = nil
	s.length = 0
}

// Empty:判断栈是否为空
func (s *Stack) Empty() bool {
	return s.length == 0
}

//返回栈的值，通过copy的方式，不会对原值造成影响
func (s *Stack) Values() []interface{} {
	ret := make([]interface{}, s.length)
	copy(ret, s.cache)
	return ret
}
