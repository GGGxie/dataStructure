package main

import (
	"fmt"
)

var counter int32

func main() {
	fmt.Println(numWays(7))
}

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/solution/go-gun-dong-shu-zu-by-xilepeng-x8wj
// 青蛙跳台阶问题
// 动态规划
func numWays(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n+1) //记录跳到i级台阶有几种方法
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007 //加法分配律？
	}
	return dp[n]
}

// // //优化方案，空间为O(1)
// func numWays(n int) int {
// 	prev, curr := 1, 1
// 	for i := 2; i <= n; i++ {
// 		next := (prev + curr) % 1000000007
// 		prev = curr
// 		curr = next
// 	}
// 	return curr
// }
