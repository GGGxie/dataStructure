package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println(uuid.NewV1())
}

// https://leetcode-cn.com/problems/single-number/
// 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 思路:数组所有元素进行异或即可
// 1. 0^A=A
// 2. A^A=A
// 3. 异或运算满足交换律和结合律  a^b^a=b^a^a=b^(a^a)=b^0=b。
func singleNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	return ret
}
