package main

import (
	"fmt"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	fmt.Println(titleToNumber("AB"))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xa6dkt/
// 26进制转10进制
func titleToNumber(columnTitle string) int {
	var ret int
	length := len(columnTitle)
	for i, ch := range columnTitle {
		temp := exponent(26, (length - i - 1)) //位数
		temp2 := (int(ch-'A') + 1)             //最高位
		ret += temp2 * temp
	}
	return ret
}

//计算a的n次方
func exponent(a, n int) int {
	result := int(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}
