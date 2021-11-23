package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
}

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
// 左旋转字符串
func reverseLeftWords(s string, n int) string {
	length := len(s)
	n %= length //防止n>length
	return s[n:] + s[0:n]
}
