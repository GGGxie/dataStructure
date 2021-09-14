package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}

var mapp [][]int //记录点的最长递增路径值
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

	maxNum := math.MinInt32
	for i := range matrix {
		for j := range matrix[i] {
			maxNum = max(dfs(i, j, matrix), maxNum)
		}
	}
	return maxNum
}

func dfs(row, col int, matrix [][]int) int {
	if mapp[row][col] != 0 {
		return mapp[row][col]
	}
	mapp[row][col]++
	for _, s := range dis {
		newRow, newCol := row+s[0], col+s[1]
		if lenRow > newRow && newRow >= 0 && lenCol > newCol && newCol >= 0 && matrix[newRow][newCol] > matrix[row][col] {
			mapp[row][col] = max(dfs(newRow, newCol, matrix)+1, mapp[row][col])
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
