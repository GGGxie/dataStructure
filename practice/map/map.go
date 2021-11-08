package mapp

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2osfr/
// 矩阵中的最长递增路径,找出其中图中的最长递增路径的长度。
var mapp [][]int //记录每个点的最长递增路径值
var (
	dis    = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	lenRow int
	lenCol int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	lenRow = len(matrix)
	lenCol = len(matrix[0])
	//初始化mapp
	mapp = make([][]int, lenRow)
	for r := 0; r < lenRow; r++ {
		mapp[r] = make([]int, lenCol)
	}

	maxNum := math.MinInt32 //记录最长递增路径值
	for i := range matrix { //对所有节点进行深搜
		for j := range matrix[i] {
			maxNum = max(dfs(i, j, matrix), maxNum)
		}
	}
	return maxNum
}

func dfs(row, col int, matrix [][]int) int {
	if mapp[row][col] != 0 { //如果某个点已经被dfs过,则mapp中已经存有其长递增路劲值
		return mapp[row][col]
	}
	mapp[row][col]++
	for _, s := range dis {
		newRow, newCol := row+s[0], col+s[1]
		if lenRow > newRow && newRow >= 0 && lenCol > newCol && newCol >= 0 && matrix[newRow][newCol] > matrix[row][col] {
			mapp[row][col] = max(dfs(newRow, newCol, matrix)+1, mapp[row][col]) //递归深搜
		}
	}
	return mapp[row][col]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaorig/
// 单词搜索 II
// 暴力会超时
// func findWords(board [][]byte, words []string) []string {
// 	if board == nil {
// 		return nil
// 	}
// 	lenRow = len(board)
// 	lenCol = len(board[0])
// 	var ret []string
// 	for _, w := range words {
// 		if ok := _findWord(board, w); ok {
// 			ret = append(ret, w)
// 		}
// 	}
// 	return ret
// }
// func _findWord(board [][]byte, words string) bool {
// 	flag = make([][]int, lenRow)
// 	for row := range flag {
// 		flag[row] = make([]int, lenCol)
// 	}
// 	for row := range board {
// 		for col := range board[row] {
// 			if board[row][col] == words[0] {
// 				if ok := dfs(board, row, col, words, 0); ok {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }
// func dfs(board [][]byte, row, col int, word string, index int) bool { //对bard[row][col]开始深度搜索，判断与word[index]是否匹配
// 	if board[row][col] == word[index] && index == len(word)-1 {
// 		return true
// 	}
// 	if board[row][col] == word[index] {
// 		flag[row][col] = 1 //标记board[row][col]已经被搜索
// 		for _, d := range dis {
// 			newRow := row + d[0]
// 			newCol := col + d[1]
// 			if lenRow > newRow && newRow >= 0 && lenCol > newCol && newCol >= 0 && flag[newRow][newCol] != 1 {
// 				if ok := dfs(board, newRow, newCol, word, index+1); ok {
// 					return true
// 				}
// 			}
// 		}
// 		flag[row][col] = 0 //取消标记
// 	}
// 	return false
// }

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/
// 方法一,二分,还没看懂
// 方法二:暴力法,找到给出矩阵中第k个最小的值,把map转化成slice,然后排序
// 时间复杂度: O(n^2\logn),对 n^2n个数排序
// 空间复杂度: O(n^2),一维数组需要存储这 n^2个元素
func kthSmallest(matrix [][]int, k int) int {
	var s sort.IntSlice
	for i := range matrix {
		for j := range matrix[i] {
			s = append(s, matrix[i][j])
		}
	}
	sort.Sort(s)
	return s[k-1]
}

// https://leetcode-cn.com/problems/word-ladder/
// 单词接龙
// bfs解决最短路径问题
// dfs不能解决最短路径问题
type node struct {
	val   string
	count int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var mapp = make(map[string]bool) //标记单词是否已经遍历过
	nodeList := make([]node, 0, len(wordList))
	nodeList = append(nodeList, node{
		val:   beginWord,
		count: 1,
	})
	mapp[beginWord] = true
	for {
		if len(nodeList) == 0 {
			break
		}
		top := nodeList[0]
		nodeList = nodeList[1:]
		for _, str := range wordList { //把beginWord可以到达的str压入队列
			if judge(top.val, str) && !mapp[str] {
				if str == endWord {
					return top.count + 1
				}
				nodeList = append(nodeList, node{
					val:   str,
					count: top.count + 1,
				})
				mapp[str] = true
			}
		}
	}
	return 0
}

func judge(str1, str2 string) bool { //判断两个字符串是否相差一个字符
	var flag bool
	len1 := len(str1)
	len2 := len(str2)
	if len1 != len2 {
		return false
	}
	for i := 0; i < len1; i++ {
		if str1[i] != str2[i] {
			if flag {
				return false
			} else {
				flag = true
			}
		}
	}
	return true
}
