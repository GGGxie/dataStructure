package search

// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
// 二分搜索for循环版本
func isBadVersion(version int) bool
func firstBadVersion(n int) int {
	var start, end int
	for start, end = 1, n; start < end; { //二分搜索
		mid := (start + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

// 二分搜索递归版本
// 时间复杂度O(logn),空间复杂度O(1)
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

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
// 第一个只出现一次的字符
func firstUniqChar(s string) byte {
	mapp := make(map[rune]int) //记录每个字符出现的次数
	for _, r := range s {      //遍历记录次数
		mapp[r]++
	}
	for _, r := range s { //遍历找出次数为1的字符
		if mapp[r] == 1 {
			return byte(r)
		}
	}
	return ' '
}
