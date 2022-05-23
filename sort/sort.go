package sort

import "testing"

func F2(t *testing.T) {

}

// // 给你两个有序整数数组nums1 和 nums2，请你将 nums2 合并到nums1中，使 nums1 成为一个有序数组。
// // 初始化nums1 和 nums2 的元素数量分别为m 和 n 。你可以假设nums1 的空间大小等于m + n，这样它就有足够的空间保存来自 nums2 的元素。
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	index1, index2, end := m-1, n-1, m+n-1
// 	for  index1 >= 0 && index2 >= 0 end-- {
// 		if nums1[index1] > nums2[index2] {
// 			nums1[end] = nums1[index1]
// 			index1--
// 		} else {
// 			nums1[end] = nums2[index2]
// 			index2--
// 		}
// 	}
// 	for index1 >= 0 {
// 		nums1[end] = nums1[index1]
// 		index1--
// 		end--
// 	}
// 	for index2 >= 0 {
// 		nums1[end] = nums2[index2]
// 		index2--
// 		end--
// 	}
// }

//快速排序
//时间复杂度O(nlogn), 空间复杂度O(logn)，不稳定排序
func QuickSort(slice []int) {
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

// 归并排序
// 时间复杂度O(nlogn)，空间复杂度：O(n)，稳定排序
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right)) //递归调用
}

// 把两个有序数据按从小到大进行合并，得到新的有序数组
func merge(left []int, right []int) []int {
	leftNum, rightNum := len(left), len(right)
	leftIndex, rightIndex := 0, 0
	result := make([]int, 0, leftNum+rightNum)
	for leftIndex < leftNum && rightIndex < rightNum {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	for leftIndex < leftNum {
		result = append(result, left[leftIndex])
		leftIndex++
	}
	for rightIndex < rightNum {
		result = append(result, right[rightIndex])
		rightIndex++
	}
	return result
}

//计数排序,max要比数组里面所有的元素大
// 时间复杂度O(nlogn)，空间复杂度：O(n)，稳定排序
func countingSort(array []int, max int) {
	//b中间数据，记录array排序后的顺序
	b, t := make([]int, len(array)), make([]int, max)
	for i := 0; i < len(array); i++ {
		t[array[i]]++
	}
	// fmt.Println(t)
	//实现稳定性
	for i := 1; i < max; i++ {
		t[i] += t[i-1]
	}
	// fmt.Println(t)
	//倒序实现稳定性，使得排序后的数组也是倒着来
	for j := len(array) - 1; j >= 0; j-- {
		b[t[array[j]]-1] = array[j]
		// fmt.Println(array[j], t[array[j]], t[array[j]]-1)
		t[array[j]]--
	}
	// fmt.Println(b)
	//赋值
	for i := 0; i < len(array); i++ {
		array[i] = b[i]
	}
}
