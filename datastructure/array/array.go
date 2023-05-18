package array

type Array struct {
	elements []interface{}
	size     int
}

// 判断数组中的元素是否唯一
// true:唯一
// false:不唯一
func (a *Array) Isunique() bool {
	mapp := make(map[interface{}]int)
	for _, ch := range a.elements {
		mapp[ch]++
		if mapp[ch] > 1 {
			return false
		}
	}
	return true
}

// 给定一个整数数组，找出总和最大的连续数列，并返回总和。
// 动态规划
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
