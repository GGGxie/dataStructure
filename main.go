package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mapp := make(map[int]bool)
	max := math.MinInt32
	for i := range nums {
		mapp[nums[i]] = true
	}
	for i := range nums {
		if !mapp[nums[i]-1] { //起点
			count := 1
			start := nums[i]
			for mapp[start+1] {
				count++
				start += 1
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
