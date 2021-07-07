package main

// 给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
// 初始化nums1 和 nums2 的元素数量分别为m 和 n 。你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2, end := m-1, n-1, m+n-1
	for ; index1 >= 0 && index2 >= 0; end-- {
		if nums1[index1] > nums2[index2] {
			nums1[end] = nums1[index1]
			index1--
		} else {
			nums1[end] = nums2[index2]
			index2--
		}
	}
	for index1 >= 0 {
		nums1[end] = nums1[index1]
		index1--
		end--
	}
	for index2 >= 0 {
		nums1[end] = nums2[index2]
		index2--
		end--
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
		index := left + 1 //标记比slice[flag]大的位置
		for i := index; i <= right; i++ {
			if slice[i] < slice[flag] {
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
