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

// // https://leetcode-cn.com/leetbook/read/top-interview-questions/x2p3cd/
// // 岛屿数量
// // 遍历图，对岛屿进行深搜，搜到的都标记为1
// var dir = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} //前、后、左、右

// //深搜，都记录为1
// func numIslands(grid [][]byte) int {
// 	var count int
// 	var mapp [][]bool //标记
// 	mapp = make([][]bool, len(grid))
// 	for i := range mapp {
// 		mapp[i] = make([]bool, len(grid[0]))
// 	}
// 	for x := range grid { //对图进行遍历
// 		for y := range grid[x] {
// 			if grid[x][y] == '1' && !mapp[x][y] { //对没有标记过的岛屿进行深搜
// 				dfs(x, y, grid, mapp)
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func dfs(x, y int, grid [][]byte, mapp [][]bool) {
// 	for _, d := range dir {
// 		newX := d[0] + x
// 		newY := d[1] + y
// 		if (newX < len(grid) && newX >= 0) && (newY < len(grid[0]) && newY >= 0) && grid[x][y] == '1' && !mapp[newX][newY] {
// 			mapp[newX][newY] = true //搜到的岛屿都标记为true
// 			dfs(newX, newY, grid, mapp)
// 		}
// 	}
// }

// https://leetcode-cn.com/problems/course-schedule/
// 课程表
// 判断是否有向无环图
// 1.找出所有入度为0的加入队列list中
// 2.进行bfs
// 3.遍历到的节点出线对应的节点入度-1
// 4.再找出所有入度为0的加入队列list中
// 5.判断队列list的节点数和numCourse是否相同
func canFinish(numCourses int, prerequisites [][]int) bool {
	var clsList []int                  //记录入度为0的节点
	list := make([]int, 0, numCourses) //bfs队列
	preNum := make([]int, numCourses)  //存储入度
	visit := make([]bool, numCourses)  //记录节点是否已被记录
	for i := range prerequisites {
		preNum[prerequisites[i][0]]++
	}
	//入度为0的就加入队列中
	for i := range preNum {
		if preNum[i] == 0 {
			list = append(list, i)
			clsList = append(clsList, i)
			visit[i] = true
		}
	}
	//bfs
	for len(list) != 0 {
		clsNum := list[0]
		list = list[1:]
		for i := range prerequisites { //找到所有出线，对应的节点入度-1
			if prerequisites[i][1] == clsNum {
				preNum[prerequisites[i][0]]--
			}
		}
		//bfs，入度为0的就加入队列中
		for i := range preNum {
			if preNum[i] == 0 && !visit[i] {
				list = append(list, i)
				clsList = append(clsList, i)
				visit[i] = true
			}
		}
	}
	return len(clsList) == numCourses
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2a743/
// 课程表 II
func findOrder(numCourses int, prerequisites [][]int) []int {
	var clsList []int                  //记录入度为0的节点
	list := make([]int, 0, numCourses) //bfs队列
	preNum := make([]int, numCourses)  //存储入度
	visit := make([]bool, numCourses)  //记录节点是否已被记录
	for i := range prerequisites {
		preNum[prerequisites[i][0]]++
	}
	//入度为0的就加入队列中
	for i := range preNum {
		if preNum[i] == 0 {
			list = append(list, i)
			clsList = append(clsList, i)
			visit[i] = true
		}
	}
	//bfs
	for len(list) != 0 {
		clsNum := list[0]
		list = list[1:]
		for i := range prerequisites { //找到所有出线，对应的节点入度-1
			if prerequisites[i][1] == clsNum {
				preNum[prerequisites[i][0]]--
			}
		}
		//bfs，入度为0的就加入队列中
		for i := range preNum {
			if preNum[i] == 0 && !visit[i] {
				list = append(list, i)
				clsList = append(clsList, i)
				visit[i] = true
			}
		}
	}
	if len(clsList) == numCourses {
		return clsList
	} else {
		return nil
	}
}

// // https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// // 矩阵中的路径
// // 待优化
// var (
// 	dir = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
// )

// func exist(board [][]byte, word string) bool {
// 	for row := range board {
// 		for col := range board[row] { //遍历二维切片
// 			if board[row][col] == word[0] { //对与首字母相等的进行深搜
// 				var mapp [][]bool //标记
// 				mapp = make([][]bool, len(board))
// 				for i := range mapp {
// 					mapp[i] = make([]bool, len(board[0]))
// 				}
// 				if ok := dfs(row, col, board, mapp, word, 1); ok {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }
// func dfs(row, col int, grid [][]byte, mapp [][]bool, target string, count int) bool {
// 	mapp[row][col] = true
// 	if count == len(target) {
// 		return true
// 	}
// 	for _, d := range dir {
// 		newRow := d[0] + row
// 		newCol := d[1] + col
// 		//判断newRow、newCol是否超出范围，判断grid[newRow][newCol]与下一个字母是否相等，判断grid[newRow][newCol]是否被访问过
// 		if (newRow < len(grid) && newRow >= 0) && (newCol < len(grid[0]) && newCol >= 0) && grid[newRow][newCol] == target[count] && !mapp[newRow][newCol] {
// 			if ok := dfs(newRow, newCol, grid, mapp, target, count+1); ok {
// 				return true
// 			}
// 			mapp[newRow][newCol] = false
// 		}
// 	}
// 	return false
// }
