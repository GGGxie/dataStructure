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
