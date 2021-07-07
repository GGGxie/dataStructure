package main

import "fmt"

func main() {
	a := []int{55, 123, 56, 4, 14, 56, 77}
	quickSort(a)
	fmt.Println(a)
}

//快速排序
//时间复杂度O(nlogn), 空间复杂度O(logn),不稳定排序
func quickSort(slice []int) {
	var (
		_quickSort func(left, right int, slice []int)     //利用递归不断对分区进行排序
		partition  func(left, right int, slice []int) int //排序
	)
	partition = func(left, right int, slice []int) int {
		flag := left      //基准
		index := left + 1 //标记比slice[flag]大的位置
		for i := index; i <= right; i++ {
			if slice[i] < slice[flag] {
				slice[i], slice[index] = slice[index], slice[i]
				index++
			}
		}
		slice[flag], slice[index-1] = slice[index-1], slice[flag]
		return (index - 1)
	}
	_quickSort = func(left, right int, slice []int) {
		if left < right {
			partitionIndex := partition(left, right, slice) //排序并获取基准位置
			//以基准位置进行分区，进行再排序
			_quickSort(left, partitionIndex-1, slice)
			_quickSort(partitionIndex+1, right, slice)
		}
	}
	left, right := 0, len(slice)-1 //left起始值下标，right末尾值下标
	_quickSort(left, right, slice)
}

// func waysToChange(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	dp := make([]int, n+1)
// 	dp[0] = 1
// 	for i := 1; i <= n; i++ {
// 		if i >= 1 {
// 			dp[i] += dp[i-1]
// 		}
// 		if i >= 5 {
// 			dp[i] += dp[i-5]
// 		}
// 		if i >= 10 {
// 			dp[i] += dp[i-10]
// 		}
// 		if i >= 25 {
// 			dp[i] += dp[i-25]
// 		}
// 		fmt.Println(i, dp[i])
// 	}
// 	return dp[n]
// }

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
