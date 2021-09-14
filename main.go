package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2echt/
// 零钱兑换，最少的硬币个数 。如果没有任何一种硬币组合能组成总金额
// 动态转移方程 dp[i]=min(dp[i-coins[j]]+1),其中coins[j] <= i
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp[i]：i代表amount为i时，需要的最少硬币数
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minNum := math.MaxInt32
		for j := range coins {
			if coins[j] <= i {
				minNum = min(minNum, dp[i-coins[j]]+1)
			}
		}
		dp[i] = minNum
	}
	if dp[amount] == math.MaxInt32 { //没有匹配返回-1
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
