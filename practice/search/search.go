package search

import "math"

// https://leetcode-cn.com/problems/find-peak-element/
// 寻找峰值
// 如果 {nums}[i-1] < {nums}[i] > {nums}[i+1]nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，我们可以直接返回 i 作为答案；
// 如果 {nums}[i-1] < {nums}[i] < {nums}[i+1]nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，我们需要往右走，即 i←i+1；
// 如果 {nums}[i-1] > {nums}[i] > {nums}[i+1]nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，我们需要往左走，即 i←i−1；
// 如果 {nums}[i-1] > {nums}[i] < {nums}[i+1]nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，我们可以朝任意方向走。
// 在这个基础上，再用二分查找
func findPeakElement(nums []int) int {
	length := len(nums)
	var get func(int) int
	get = func(inx int) int {
		if inx >= length || inx < 0 {
			return math.MinInt64 //超出了界限，则返回最小值
		}
		return nums[inx]
	}
	left, right := 0, length
	for {
		mid := (left + right) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		} else if get(mid) > get(mid-1) {
			left = mid + 1
			continue
		} else {
			right = mid - 1
			continue
		}
	}
}
