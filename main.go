// main.go
package main

import (
	"fmt"
	"math"
	_ "net/http/pprof"
)

var datas []string

func main() {
	mat := [][]int{{4, 5, 6}, {2, 3, 4}, {2, 3, 4}}
	fmt.Println(smallestCommonElement(mat))
}

// 找出所有行中最小公共元素
// https://leetcode.cn/problems/find-smallest-common-element-in-all-rows/
// 不能用 map,因为 map 是无序的,不能保证先遍历到最小公共元素,例如:
// mat = [[1,2,3],[2,3,4],[2,3,5]]
// 如果先遍历到 3 就错了
func smallestCommonElement(mat [][]int) int {
	mapp := map[int]int{}
	count, length := math.MaxInt, len(mat)
	for _, s := range mat {
		for j := range s {
			mapp[s[j]]++
		}
	}
	for i := range mapp {
		fmt.Println(i)
		if mapp[i] >= length && count > mapp[i] {
			count = i
		}
	}
	if count == math.MaxInt {
		count = -1
	}
	return count
}
