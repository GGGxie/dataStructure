package main

import "fmt"

func main() {
	m := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(pathWithObstacles(m))
}
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var (
		result [][]int
		rowEnd = len(obstacleGrid) - 1
		colEnd = len(obstacleGrid[0]) - 1
		dfs    func([][]int) //内部函数递归需要用到该变量
	)
	dfs = func(path [][]int) {
		if len(result) != 0 { //已找到路径
			return
		}
		row, col := path[len(path)-1][0], path[len(path)-1][1]
		if obstacleGrid[row][col] == 1 { //判断路径是否可通
			return
		}
		if row == rowEnd && col == colEnd {
			result = make([][]int, rowEnd+colEnd+1) //结果的长度必定是长加高-1
			fmt.Println(len(result))
			copy(result, path)
		}
		obstacleGrid[row][col] = 1
		if row < rowEnd { //向下走
			dfs(append(path, []int{row + 1, col})) //采用append的方式可以不改变path的值，让path仍然记录当前位置
		}
		if col < colEnd { //向右走
			dfs(append(path, []int{row, col + 1}))
		}
	}
	dfs([][]int{{0, 0}})
	return result
}

// func findClosedNumbers(num int) []int {
// 	count := func(num int) int { //找出1的总数
// 		var sum int
// 		for num != 0 {
// 			if num&1 == 1 {
// 				sum++
// 			}
// 			num >>= 1
// 		}
// 		return sum
// 	}

// 	//找到最大的值
// 	larger, smaller := -1, -1
// 	sigA := 0b01
// 	sigB := 0b10
// 	var index int
// 	for index = 0; index <= 30; index++ { //从右往左找到第一个01的位置
// 		if num&(sigA<<index) == sigA {
// 			temp := (num & ((1 << index) - 1)) //获取
// 			num = ((sigB << index) | temp)
// 			break
// 		}
// 	}
// 	//获取偏大值
// 	for ; index >= 0; index-- {

// 	}
// 	//获取偏小值
// 	temp = num - 1
// 	for temp >= 1 {
// 		if count(temp) == sumOfNum {
// 			smaller = temp
// 			break
// 		}
// 		temp--
// 	}
// 	return []int{larger, smaller}
// }
