package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
// 寻找0～n-1中缺失的数字
func missingNumber(nums []int) int {
	length := len(nums)
	for i := 0; i <= length; i++ {
		if i == length || nums[i] != i {
			return i
		}
	}
	return -1
}
