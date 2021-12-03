package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(translateNum(18822))
}

// https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
// 把数字翻译成字符串
// 动态规划，可以用滚动数组优化，待优化
// dp[i]=dp[i-2]+dp[i-1] src[i-1:i+1] <= "25" && src[i-1:i+1] >= "10"
// dp[i]=dp[i-1]//num[i] src[i-1:i+1] > "25" && src[i-1:i+1] < "10"
func translateNum(num int) int {
	src := strconv.Itoa(num)
	dp := make([]int, len(src))
	dp[0] = 1
	for i := 1; i < len(src); i++ {
		if src[i-1:i+1] <= "25" && src[i-1:i+1] >= "10" { //src[i-1:i+1]作为一个整体
			if i == 1 {
				dp[i] = dp[0] + 1
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else { //src[i]作为一个单体
			dp[i] = dp[i-1]
		}
	}
	return dp[len(src)-1]
}
