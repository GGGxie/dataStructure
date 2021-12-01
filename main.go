package main

import (
	"fmt"
)

var counter int32

func main() {
	fmt.Println(fib(45))
}

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	temp1 := 0 //指向f(n-2)
	temp2 := 1 //指向f(n-1)
	for i := 2; i <= n; i++ {
		temp3 := (temp1 + temp2) % 1000000007 //计算出f(n)
		temp1 = temp2                         //f(n-2)指向之前f(n-1)的位置
		temp2 = temp3                         //f(n-1)指向之前f(n)的位置
	}
	return temp2
}
