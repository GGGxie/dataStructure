package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseList(nil))
}

// https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
// 替换空格
func replaceSpace(s string) string {
	// return strings.ReplaceAll(s, " ", "%20")
	var ret string
	for _, r := range s {
		if r == ' ' {
			ret += "%20"
		} else {
			ret += string(r)
		}
	}
	return ret
}
