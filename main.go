package main

import (
	_ "net/http/pprof"
)

func main() {}

// https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
// 礼物的最大价值
// 动态规划
// 动态转移方程:maxValue(x, y) = max(maxValue(x-1,y)+grid[x][y],maxValue(x,y-1)+grid[x][y])
func maxValue(grid [][]int) int {
	var dp [][]int //记录和grid对应的每个点的最大价值
	dp = make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[i]))
	}
	for row := range grid {
		for col := range grid[row] {
			if row == 0 && col == 0 { //grid[0][0]
				dp[0][0] = grid[0][0]
			} else if row == 0 { //第一行
				dp[row][col] = dp[row][col-1] + grid[row][col]
			} else if col == 0 { //第一列
				dp[row][col] = dp[row-1][col] + grid[row][col]
			} else {
				dp[row][col] = max(dp[row][col-1], dp[row-1][col]) + grid[row][col]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
