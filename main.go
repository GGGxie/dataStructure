package main

import (
	"fmt"
)

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
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
	fmt.Println(mapp)
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
