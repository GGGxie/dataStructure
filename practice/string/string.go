package string

import (
	"strings"
)

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xapbdt/
//字符串翻转
func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

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

//https://leetcode-cn.com/leetbook/read/top-interview-questions/xaorig/
//单词搜索 II,给出一个图和切片字符串，判断切片字符串中的字符串是否在图中
//方法一：对图中每一个字母进行深搜，会超时
var (
	flag   [][]int
	dis    = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	lenCol int
	lenRow int
)

func findWords(board [][]byte, words []string) []string {
	if board == nil {
		return nil
	}
	lenRow = len(board)
	lenCol = len(board[0])
	var ret []string
	for _, w := range words {
		if ok := _findWord(board, w); ok {
			ret = append(ret, w)
		}
	}
	return ret
}
func _findWord(board [][]byte, words string) bool {
	flag = make([][]int, lenRow)
	for row := range flag {
		flag[row] = make([]int, lenCol)
	}
	for row := range board {
		for col := range board[row] {
			if board[row][col] == words[0] {
				if ok := dfs(board, row, col, words, 0); ok {
					return true
				}
			}
		}
	}
	return false
}
func dfs(board [][]byte, row, col int, word string, index int) bool { //对bard[row][col]开始深度搜索，判断与word[index]是否匹配
	if board[row][col] == word[index] && index == len(word)-1 {
		return true
	}
	if board[row][col] == word[index] {
		flag[row][col] = 1 //标记board[row][col]已经被搜索
		for _, d := range dis {
			newRow := row + d[0]
			newCol := col + d[1]
			if lenRow > newRow && newRow >= 0 && lenCol > newCol && newCol >= 0 && flag[newRow][newCol] != 1 {
				if ok := dfs(board, newRow, newCol, word, index+1); ok {
					return true
				}
			}
		}
		flag[row][col] = 0 //取消标记
	}
	return false
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

// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
// 左旋转字符串
func reverseLeftWords(s string, n int) string {
	length := len(s)
	n %= length //防止n>length
	return s[n:] + s[0:n]
}
