package main

import (
	"fmt"
)

var counter int32

func main() {
	a := []int{7, 6}
	fmt.Println(maxProfit1(a))
}

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/submissions/
// 连续子数组的最大和
// maxSubArray(n)=max(maxSubArray(n-1)+nums[n],nums[n])
// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	ret := dp[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], dp[i-1]+nums[i])
// 		if dp[i] > ret {
// 			ret = dp[i]
// 		}
// 	}
// 	return ret
// }
// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
// 连续子数组的最大和
// 动态规划: maxSubArray(n)=max(maxSubArray(n-1)+nums[n],nums[n])
// 用滚动数组优化空间
func maxSubArray(nums []int) int {
	temp, ret := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		temp2 := max(nums[i], temp+nums[i]) //temp:记录maxSubArray(i-1)的最大值,temp2:记录maxSubArray(i)的最大值
		if temp2 > ret {
			ret = temp2
		}
		temp = temp2 //滚动,temp记录maxSubArray(i)的最大值,继续遍历i+1
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
