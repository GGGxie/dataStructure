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

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xm42hs/
// 旋转数组,给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	length := len(nums)
	k %= length //当k>length
	index := length - k
	copy(nums, append(nums[index:], nums[0:index]...))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xm1rfd/
// 存在重复元素.给定一个整数数组，判断是否存在重复元素.
func containsDuplicate(nums []int) bool {
	mapp := make(map[int]int)
	for i := range nums {
		mapp[nums[i]]++
		if mapp[nums[i]] > 1 {
			return true
		}
	}
	return false
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmy9jh/
// 移动零.给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 双指针实现,
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums) //left:第一个为0的下标,rigth:遍历切片的下标
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// 搜索二维矩阵 II,
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 暴力对每一行进行二分
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if binarySearch(row, target) {
			return true
		}
	}
	return false
}

//二分搜索
func binarySearch(list []int, target int) bool {
	if len(list) == 0 {
		return false
	}
	mid := len(list) / 2
	if list[mid] == target {
		return true
	}
	if list[mid] > target {
		return binarySearch(list[0:mid], target) //当len(list)=1,list[0:0]为空，len(list[0:0])=0
	} else if list[mid] < target {
		return binarySearch(list[mid+1:], target) //当len(list)=1,list[1:]为空,len(list[1:])=0
	}
	return false
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
