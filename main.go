package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}

// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
// 最长不含重复字符的子字符串
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	//滑动窗口+hash
	if len(s) == 0 {
		return 0
	}
	ret := 0
	mapp := make(map[byte]struct{}) //标记数组中的重复元素，空结构体做占位符节省空间
	for left, right := 0, 0; right < len(s); {
		if _, ok := mapp[s[right]]; !ok { //窗口右边界往右扩展一个元素
			mapp[s[right]] = struct{}{}
			right++
		} else { //窗口左边界往右收缩一个元素,一直滑动到set中没有重复的元素
			delete(mapp, s[left])
			left++
		}
		if ret < right-left {
			ret = right - left
		}
	}
	return ret
}
