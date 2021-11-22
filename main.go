package main

import (
	"fmt"
	"math"
)

func main() {
	// n := uint32(0b00000010100101000001111010011100)
	fmt.Println(reverse(-123))
}

func hammingWeight(num uint32) int {
	var ret int
	for i := 0; i < 32; i++ {
		ret += int(num & 1)
		num >>= 1
	}
	return ret
}

// https://leetcode-cn.com/problems/reverse-integer/
// 整数反转
func reverse(x int) int {
	ret := 0
	for x != 0 {
		temp := x % 10
		x /= 10
		ret *= 10
		ret += temp
	}
	if ret > int(math.MaxInt32) || ret < int(-math.MaxInt32) {
		return 0
	}
	return ret
}
