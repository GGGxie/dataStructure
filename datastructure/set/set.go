package set

import (
	"fmt"
)

// 用map实现set
type Set map[string]struct{}

// 判断集合中是否拥有某个元素
func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

// 给集合添加一个元素
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

// 删除集合中的某个元素
func (s Set) Delete(key string) {
	delete(s, key)
}

// 遍历输出集合中的元素
func (s Set) Range() {
	for i, _ := range s {
		fmt.Println(i)
	}
}
func test() {
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	s.Range()
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Jack"))
}
