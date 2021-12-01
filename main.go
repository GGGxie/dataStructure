package main

import (
	"fmt"
	"math"
)

var counter int32

func main() {
	a := []int{7, 6}
	fmt.Println(maxProfit1(a))
}

// https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/submissions/
// 股票的最大利润
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := math.MaxInt32 //维护一个最小值
	ret := 0             //最大利润
	for i := 0; i < len(prices); i++ {
		if prices[i] < min { //当但当前值比最小值大，最小值为当前值
			min = prices[i]
		}
		if prices[i]-min > ret { //如果当前值-最小值>最大利润，则最大利润=当前值-最小值
			ret = prices[i] - min
		}
	}
	return ret
}
