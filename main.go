package main

func main() {
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		flag     map[int]bool = make(map[int]bool) //标记节点时候已经被染色
		dfs      func(tempSr, tempSc int)
		rLen     = len(image)
		cLen     = len(image[0])
		oldColor = image[sr][sc] //记录原始节点颜色
	)
	dfs = func(tempSr, tempSc int) { //深度搜索
		if image[tempSr][tempSc] != oldColor {
			return
		}
		image[tempSr][tempSc] = newColor
		flag[tempSr*cLen+tempSc] = true
		if 0 <= tempSr-1 && !flag[(tempSr-1)*cLen+tempSc] { //上
			dfs(tempSr-1, tempSc)
		}
		if tempSr+1 < rLen && !flag[(tempSr+1)*cLen+tempSc] { //下
			dfs(tempSr+1, tempSc)
		}
		if 0 <= tempSc-1 && !flag[tempSr*cLen+(tempSc-1)] { //左
			dfs(tempSr, tempSc-1)
		}
		if tempSc+1 < cLen && !flag[tempSr*cLen+(tempSc+1)] { //右
			dfs(tempSr, tempSc+1)
		}
	}
	dfs(sr, sc)
	return image
}

// func multiply(A int, B int) int {
// 	var re func(A, B *int, C int)
// 	re = func(A, B *int, C int) {
// 		if *B == 0 {
// 			return
// 		}
// 		*A += C
// 		*B -= 1
// 		re(A, B, C)
// 	}
// 	if A > B {
// 		B -= 1
// 		re(&A, &B, A)
// 		return A
// 	} else {
// 		A -= 1
// 		re(&B, &A, B)
// 		return B
// 	}
// }

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
