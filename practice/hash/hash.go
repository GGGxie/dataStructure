package hash

// https://leetcode-cn.com/problems/4sum-ii/
// 四数相加
// 暴力的时间复杂度为O(n^4),用两两循环并存在hash中进行相加，时间复杂度降到O(n^2)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var ret int
	mapp := make(map[int]int)
	for _, i := range nums1 { //把nums1+nums2相加的值放入mapp
		for _, j := range nums2 {
			mapp[i+j]++
		}
	}
	for _, i := range nums3 { //看看-(nums3+nums4)有多少个存在mapp中,存在一个便有一个结果
		for _, j := range nums4 {
			ret += mapp[-(i + j)]
		}
	}
	return ret
}
