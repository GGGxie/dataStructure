package arrayStack

type Stack struct {
	cache []interface{}
	size  int
}

func (s *Stack) Push(value interface{}) {
	s.cache = append(s.cache, value)

}

func Pop(value interface{}) (ok bool) {

}

func Peek() (value interface{}, ok bool) {

}

func Empty() bool {

}

func Size() int {

}

func Clear() {

}

func Values() []interface{} {

}
