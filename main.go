package main

import "fmt"

func main() {
	a := []int{55, 60, 56, 4, 14, 56, 77, 14}
	// a := []int{3, 2}
	countingSort(a, 78)
	fmt.Println(a)
}

//计数排序
func countingSort(array []int, k int) {
	//b中间数据，记录array排序后的顺序
	b, t := make([]int, len(array)), make([]int, k)
	for i := 0; i < len(array); i++ {
		t[array[i]]++
	}
	fmt.Println(t)
	//实现稳定性
	for i := 1; i < k; i++ {
		t[i] += t[i-1]
	}
	fmt.Println(t)
	//倒序实现稳定性，使得排序后的数组也是倒着来
	for j := len(array) - 1; j >= 0; j-- {
		b[t[array[j]]-1] = array[j]
		fmt.Println(array[j], t[array[j]], t[array[j]]-1)
		t[array[j]]--
	}
	fmt.Println(b)
	//赋值
	for i := 0; i < len(array); i++ {
		array[i] = b[i]
	}
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
