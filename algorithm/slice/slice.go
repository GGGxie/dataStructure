package slice

// 找出所有行中最小公共元素
// https://leetcode.cn/problems/find-smallest-common-element-in-all-rows/
// 不能用 map,因为 map 是无序的,不能保证先遍历到最小公共元素,例如:
// mat = [[1,2,3],[2,3,4],[2,3,5]]
// 如果先遍历到 3 就错了
func smallestCommonElement(mat [][]int) int {
	slice := [10001]int{}
	length := len(mat)
	for _, s := range mat {
		for j := range s {
			slice[s[j]]++
		}
	}
	for i := range slice {
		if slice[i] >= length {
			return i
		}
	}

	return -1
}

// 缺失的区间
// https://leetcode.cn/problems/missing-ranges/
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	ret := [][]int{}
	if len(nums) == 0 {
		ret = append(ret, []int{lower, upper})
		return ret
	}
	// 处理首位
	if nums[0]-lower >= 1 {
		ret = append(ret, []int{lower, nums[0] - 1})
	}
	// 处理 nums 列表
	idx := nums[0]
	for i := range nums {
		if nums[i]-idx > 1 {
			ret = append(ret, []int{idx + 1, nums[i] - 1})
		}
		idx = nums[i]
	}
	// 处理末位
	if upper-nums[len(nums)-1] >= 1 {
		ret = append(ret, []int{nums[len(nums)-1] + 1, upper})
	}
	return ret
}

// 会议室
// https://leetcode.cn/problems/meeting-rooms/
// 先排序再判断后一场会议的时间是否小于前一场会议的结束时间
func canAttendMeetings(intervals [][]int) bool {
	QuickSort := func(slice [][]int) {
		var (
			_quickSort func(left, right int, slice [][]int)     //利用递归不断对分区进行排序
			partition  func(left, right int, slice [][]int) int //排序
		)
		partition = func(left, right int, slice [][]int) int {
			flag := left      //基准
			index := left + 1 //标记比slice[flag]大的位置
			for i := index; i <= right; i++ {
				if slice[i][0] < slice[flag][0] {
					slice[i], slice[index] = slice[index], slice[i]
					index++
				}
			}
			slice[flag], slice[index-1] = slice[index-1], slice[flag]
			return (index - 1)
		}
		_quickSort = func(left, right int, slice [][]int) {
			if left < right {
				partitionIndex := partition(left, right, slice) //排序并获取基准位置
				//以基准位置进行分区，进行再排序
				_quickSort(left, partitionIndex-1, slice)
				_quickSort(partitionIndex+1, right, slice)
			}
		}
		left, right := 0, len(slice)-1 //left起始值下标，right末尾值下标
		_quickSort(left, right, slice)
	}
	QuickSort(intervals)
	for i := 1; i < len(intervals); i++ {
		// [0,1][1,2]也返回 true
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
