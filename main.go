package main

import (
	"fmt"
	"sort"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}

// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 寻找重复数
// 应该用快满指针，但是没搞懂
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
