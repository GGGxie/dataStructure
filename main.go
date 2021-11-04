package main

import (
	"fmt"
	"sort"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	a := []int{1, 3, 2, 2, 3}
	wiggleSort(a)
	fmt.Println(a)
}

// 给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

// 你可以假设所有输入数组都可以得到满足题目要求的结果。

//

// 示例 1：

// 输入：nums = [1,5,1,1,6,4]
// 输出：[1,6,1,5,1,4]
// 解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。
// 示例 2：

// 输入：nums = [1,3,2,2,3,1]
// 输出：[2,3,1,3,1,2]

//先排序，再逆序穿插
// func wiggleSort(nums []int) {
// 	l := len(nums)
// 	j, k := l, (l+1)>>1//
// 	t := make([]int, l)
// 	copy(t, nums)
// 	sort.Ints(t)
// 	fmt.Println(t)
// 	fmt.Println(j, k)
// 	for i := 0; i < l; i++ {
// 		if i&1 == 1 {
// 			j--
// 			nums[i] = t[j]
// 		} else {
// 			k--
// 			nums[i] = t[k]
// 		}
// 	}
// }

//先排序，再逆序穿插
//逆序穿插而不是正序穿插是因为逆序能把前半段和后半段相同的分开来
//逆序：[1,2,2,3]排序后是[1,2,2,3],前半段是[1,2]，后半段是[2,3],穿插后是[2,3,1,2]
//正序：[1,2,2,3]排序后是[1,2,2,3],前半段是[1,2]，后半段是[2,3],穿插后是[1,2,2,3](不满足要求)
func wiggleSort(nums []int) {
	length := len(nums)
	tempNums := make([]int, length)
	copy(tempNums, nums)
	sort.Ints(tempNums)                 //从小到大排序
	idx1, idx2 := (length+1)>>1, length //idx1:指向前半段数组的尾部，inx2指向后半段数组的尾部
	for i := 0; i < length; i++ {
		if i&1 == 1 { //偶数位
			idx2--
			nums[i] = tempNums[idx2]
		} else { //奇数位,基数时，前半段会多一个，所以必须要先插前半段
			idx1--
			nums[i] = tempNums[idx1]
		}
	}
}
