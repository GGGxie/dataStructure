package main

import (
	"fmt"

	"github.com/GGGxie/dataStructure/lists"
)

func main() {
	list := lists.InitList()
	list.Add(1)
	list.Add(3)
	list.Add(2)
	list.Insert(7, 2)
	list.Append(4)
	list.Append(5)
	list.Delete(2)
	list.Iterate()
	fmt.Println("`````````````")
	list.Reverse()
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
