package mapp

import (
	"math"
)

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2osfr/
// 矩阵中的最长递增路径,找出其中图中的最长递增路径的长度。
var mapp [][]int //记录每个点的最长递增路径值
var (
	dis    = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	lenRow int
	lenCol int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	lenRow = len(matrix)
	lenCol = len(matrix[0])
	//初始化mapp
	mapp = make([][]int, lenRow)
	for r := 0; r < lenRow; r++ {
		mapp[r] = make([]int, lenCol)
	}

	maxNum := math.MinInt32 //记录最长递增路径值
	for i := range matrix { //对所有节点进行深搜
		for j := range matrix[i] {
			maxNum = max(dfs(i, j, matrix), maxNum)
		}
	}
	return maxNum
}

func dfs(row, col int, matrix [][]int) int {
	if mapp[row][col] != 0 { //如果某个点已经被dfs过,则mapp中已经存有其长递增路劲值
		return mapp[row][col]
	}
	mapp[row][col]++
	for _, s := range dis {
		newRow, newCol := row+s[0], col+s[1]
		if lenRow > newRow && newRow >= 0 && lenCol > newCol && newCol >= 0 && matrix[newRow][newCol] > matrix[row][col] {
			mapp[row][col] = max(dfs(newRow, newCol, matrix)+1, mapp[row][col]) //递归深搜
		}
	}
	return mapp[row][col]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
