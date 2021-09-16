package string

import (
	"strings"
)

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

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xa503c/
// 单词拆分。给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
// 状态转移方程：dp[i]=dp[j] && check(s[j..i−1])
func wordBreak(s string, wordDict []string) bool {
	length := len(s)
	mapp := make(map[string]bool)
	for _, v := range wordDict {
		mapp[v] = true
	}
	dp := make([]bool, length+1) //dp[i]记录子串s[0:i]是否能在字段中被拆分成一个或多个单词
	dp[0] = true
	for i := 1; i <= length; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mapp[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[length]
}
