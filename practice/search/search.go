package search

import (
	"math"
	"sort"
)

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

// https://leetcode-cn.com/problems/find-the-duplicate-number/
// 寻找重复数
// 应该用快慢指针，但是没搞懂
func findDuplicate(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// 数组中重复的数字
func findRepeatNumber(nums []int) int {
	mapp := make(map[int]int) //存储元素的个数
	for i := range nums {
		mapp[nums[i]]++
		if mapp[nums[i]] > 1 { //大于一个元素就直接返回
			return nums[i]
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
// 在排序数组中查找数字 I
func search(nums []int, target int) int {
	var count int
	for i := range nums {
		if nums[i] == target {
			count++
		}
	}
	return count
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

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/submissions/
// 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := range matrix {
		//二分搜索
		start, end := 0, len(matrix[i])-1
		for start <= end {
			mid := (start + end) / 2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				end = mid - 1
			} else if matrix[i][mid] < target {
				start = mid + 1
			}
		}
	}
	return false
}

// https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
// 旋转数组的最小数字
func minArray(numbers []int) int {
	for i := range numbers { //处理最小元素在中间或结尾的情况
		if i != 0 && numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0] //数组顺序没有改变，那么第一个元素为最小值
}
