package main

import "fmt"

func main() {
	count := func(num int32) (sum int) { //获取数字的二进制表示中1的个数
		for num != 0 {
			sum++
			num &= (num - 1)
		}
		return
	}
	fmt.Println(count(-0b1101))
	fmt.Printf("%0b\n",0xaaaaaaaa)
}

func convertInteger(A int, B int) int {
	count := func(num int32) (sum int) {
		fmt.Println(num)
		for num != 0 {
			sum++
			num = num & (num - 1)
		}
		return
	}
	return count(int32(A ^ B))
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
