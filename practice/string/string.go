package string

import "strings"

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xah8k6/
// 判断字符串是否为回文串
func isPalindrome(s string) bool {
	stemp := strings.ToUpper(s) //全部转大写
	length := len(s)
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		for !isLetterOrNumber(stemp[left]) { //取下一个字母或数字进行比较
			left++
			if left > length-1 {
				break
			}
		}
		for !isLetterOrNumber(stemp[right]) { //取下一个字母或数字进行比较
			right--
			if right < 0 {
				break
			}
		}
		if left < right && stemp[left] != stemp[right] {
			return false
		}
	}
	return true
}

func isLetterOrNumber(ch byte) bool { //判断是否为字母或数字
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
