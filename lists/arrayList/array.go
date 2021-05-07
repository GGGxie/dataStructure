package arrayList

type Array struct {
	elements []interface{}
	size     int
}

//判断数组中的元素是否唯一
//true:唯一
//false:不唯一
func (a *Array) Isunique() bool {
	mapp := make(map[interface{}]int)
	for _, ch := range a.elements {
		mapp[ch]++
		if mapp[ch] > 1 {
			return false
		}
	}
	return true
}
