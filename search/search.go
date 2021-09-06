package search

// 假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
// 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。
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
