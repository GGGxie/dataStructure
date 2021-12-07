package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// 最长不含重复字符的子字符串
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	//滑动窗口+hash
	if len(s) == 0 {
		return 0
	}
	ret := 0
	mapp := make(map[byte]struct{}) //标记数组中的重复元素，空结构体做占位符节省空间
	for left, right := 0, 0; right < len(s); {
		if _, ok := mapp[s[right]]; !ok { //窗口右边界往右扩展一个元素
			mapp[s[right]] = struct{}{}
			right++
		} else { //窗口左边界往右收缩一个元素,一直滑动到set中没有重复的元素
			delete(mapp, s[left])
			left++
		}
		if ret < right-left {
			ret = right - left
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
