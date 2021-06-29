package maps

// //迷路的机器人，深搜找终点
// //优化点：状态压缩，二维切片优化成普通切片(行列最大不超过100，可以以101为区间大小，给切片分段)
// func pathWithObstacles(obstacleGrid [][]int) [][]int {
// 	var (
// 		result [][]int
// 		rowEnd = len(obstacleGrid) - 1
// 		colEnd = len(obstacleGrid[0]) - 1
// 		dfs    func([][]int) //内部函数递归需要用到该变量
// 	)
// 	dfs = func(path [][]int) {
// 		if len(result) != 0 { //已找到路径
// 			return
// 		}
// 		row, col := path[len(path)-1][0], path[len(path)-1][1]
// 		if obstacleGrid[row][col] == 1 { //判断路径是否可通
// 			return
// 		}
// 		if row == rowEnd && col == colEnd {
// 			result = make([][]int, rowEnd+colEnd+1) //结果的长度必定是长加高-1
// 			copy(result, path)
// 		}
// 		obstacleGrid[row][col] = 1
// 		if row < rowEnd { //向下走
// 			dfs(append(path, []int{row + 1, col})) //采用append的方式可以不改变path的值，让path仍然记录当前位置
// 		}
// 		if col < colEnd { //向右走
// 			dfs(append(path, []int{row, col + 1}))
// 		}
// 	}
// 	dfs([][]int{{0, 0}})
// 	return result
// }

//迷路的机器人，深搜找终点
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var (
		result [][]int
		rowEnd = len(obstacleGrid) - 1
		colEnd = len(obstacleGrid[0]) - 1
		dfs    func([]int) //内部函数递归需要用到该变量
		m      int         = 101
	)
	dfs = func(path []int) {
		if len(result) != 0 { //已找到路径
			return
		}
		now := path[len(path)-1]
		row, col := now/m, now%m
		if obstacleGrid[row][col] == 1 { //判断路径是否可通
			return
		}
		if row == rowEnd && col == colEnd {
			result = make([][]int, rowEnd+colEnd+1) //结果的长度必定是长加高-1
			for t, p := range path {
				result[t] = []int{p / m, p % m}
			}
		}
		obstacleGrid[row][col] = 1
		if row < rowEnd { //向下走
			dfs(append(path, now+m)) //采用append的方式可以不改变path的值，让path仍然记录当前位置
		}
		if col < colEnd { //向右走
			dfs(append(path, now+1))
		}
	}
	dfs([]int{0})
	return result
}
