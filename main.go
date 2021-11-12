package main

import "fmt"

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	n := uint32(0b00000010100101000001111010011100)
	fmt.Printf("%032b\n", reverseBits(n))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2dx36/
// 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		temp := num & 1 //取出num最后一个二进制位
		num >>= 1       //num右移一位
		ret <<= 1       //ret左移一位
		ret += temp     //ret+取出的二进制位
	}
	return ret
}
