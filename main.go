package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	m := [][]byte{{'a', 'a'}}
	b := "aaa"
	fmt.Println(exist(m, b))
}

// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/
// 矩阵中的路径
// 待优化
var (
	dir = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
)

func exist(board [][]byte, word string) bool {
	for row := range board {
		for col := range board[row] { //遍历二维切片
			if board[row][col] == word[0] { //对与首字母相等的进行深搜
				var mapp [][]bool //标记
				mapp = make([][]bool, len(board))
				for i := range mapp {
					mapp[i] = make([]bool, len(board[0]))
				}
				if ok := dfs(row, col, board, mapp, word, 1); ok {
					return true
				}
			}
		}
	}
	return false
}
func dfs(row, col int, grid [][]byte, mapp [][]bool, target string, count int) bool {
	mapp[row][col] = true
	if count == len(target) {
		return true
	}
	for _, d := range dir {
		newRow := d[0] + row
		newCol := d[1] + col
		//判断newRow、newCol是否超出范围，判断grid[newRow][newCol]与下一个字母是否相等，判断grid[newRow][newCol]是否被访问过
		if (newRow < len(grid) && newRow >= 0) && (newCol < len(grid[0]) && newCol >= 0) && grid[newRow][newCol] == target[count] && !mapp[newRow][newCol] {
			if ok := dfs(newRow, newCol, grid, mapp, target, count+1); ok {
				return true
			}
			mapp[newRow][newCol] = false
		}
	}
	return false
}
