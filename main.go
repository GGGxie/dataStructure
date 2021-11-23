package main

import (
	"fmt"
)

func main() {
	z := [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	fmt.Println(findNumberIn2DArray(z, 5))
}

// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/submissions/
// 二维数组中的查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := range matrix {
		//二分搜索
		start, end := 0, len(matrix[i])-1
		for start <= end {
			mid := (start + end) / 2
			fmt.Println(matrix[i][mid])
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				end = mid - 1
			} else if matrix[i][mid] < target {
				start = mid + 1
			}
		}
	}
	return false
}
