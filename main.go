package main

import "fmt"

func main() {
	fmt.Printf("%0b\n", 34)
	fmt.Printf("%0b\n", 36)
	fmt.Printf("%0b\n", 33)
	a := 0b1101100
	num := a & (-a)
	fmt.Printf("%0b\n", -a)
	fmt.Printf("%0b\n", num)
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
