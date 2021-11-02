package main

import (
	"fmt"
	"math/rand"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	a := map[int]bool{2: true, 3: true, 4: true}
	for i := 0; i < 100; i++ {
		for k, _ := range a {
			fmt.Println(k)
			continue
		}
	}
}

// data 按index顺序保存insert val
// map 保存insert var 和其在 data 中 index的映射
// 1 插入时 直接append， map直接存
// 2 删除时 map查到的要删除的值 slice中交换始终删除最后一个值 把最后一个值赋到实际要删除的index上 然后更新map对应值的index

type RandomizedSet struct {
	mapp map[int]int //记录数据的下标
	data []int       //记录数据的数组
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapp: make(map[int]int),
		data: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mapp[val]; !ok {
		this.mapp[val] = len(this.data)
		this.data = append(this.data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.mapp[val]; ok {
		length := len(this.data)
		this.data[index], this.data[length-1] = this.data[length-1], this.data[index] //数组尾部和下标互换
		this.data = this.data[0 : length-1]                                           //去掉尾部
		if index != length-1 {                                                        //如果刚好移除的是尾部元素，就不用替换下标
			this.mapp[this.data[index]] = index //被替换的尾部元素改下标
		}
		delete(this.mapp, val) //去掉下标存储
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	length := len(this.data)
	random := rand.Int() % length
	return this.data[random]
}
