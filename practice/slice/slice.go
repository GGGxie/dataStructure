package slice

import "math"

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmk3rv/
// 乘积最大子数组
// 动态转移方程：
// maxDp[i] = max{maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i]}
// minDp[i] = min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i]
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxDp := make([]int, length) //maxDP[i]:以第i个元素结尾的最大子数组的乘积
	minDp := make([]int, length) //minDP[i]:以第i个元素结尾的最小子数组的乘积
	maxDp[0], minDp[0] = nums[0], nums[0]
	maxNum := nums[0]
	for i := 1; i < length; i++ {
		maxDp[i] = max(max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		minDp[i] = min(min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		maxNum = max(max(maxDp[i], minDp[i]), maxNum)
	}
	return maxNum
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmz79t/
// 多数元素。给定一个大小为 n 的数组，找到其中的多数元素。
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mapp := make(map[int]int)
	for i := range nums {
		mapp[nums[i]]++
	}
	maxNum := math.MinInt32
	index := 0
	for i := range mapp {
		if mapp[i] > maxNum {
			maxNum = mapp[i]
			index = i
		}
	}
	return index
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
