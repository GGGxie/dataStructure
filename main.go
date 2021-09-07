package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(0b11111000)
}
func largestNumber(nums []int) string {
	var ret string
	quickSort(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if ret == "" && nums[i] == 0 {
			continue
		}
		ret += strconv.FormatInt(int64(nums[i]), 10)
	}
	if ret == "" {
		return "0"
	} else {
		return ret
	}
}

func max(a, b int) bool {
	s1 := strconv.FormatInt(int64(a), 10) + strconv.FormatInt(int64(b), 10)
	s2 := strconv.FormatInt(int64(b), 10) + strconv.FormatInt(int64(a), 10)
	if s1 > s2 {
		return false
	} else {
		return true
	}
}

//快速排序
//时间复杂度O(nlogn), 空间复杂度O(logn),不稳定排序
func quickSort(slice []int) {
	var (
		_quickSort func(left, right int, slice []int)     //利用递归不断对分区进行排序
		partition  func(left, right int, slice []int) int //排序
	)
	partition = func(left, right int, slice []int) int {
		flag := left      //基准
		index := left + 1 //标记比slice[flag]大的第一个位置
		for i := index; i <= right; i++ {
			if max(slice[i], slice[flag]) {
				slice[i], slice[index] = slice[index], slice[i]
				index++
			}
		}
		slice[flag], slice[index-1] = slice[index-1], slice[flag]
		return (index - 1)
	}
	_quickSort = func(left, right int, slice []int) {
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
