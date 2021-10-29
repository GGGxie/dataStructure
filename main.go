package main

import "fmt"

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	s := "32-1-1"
	fmt.Println(calculate(s))
}

// https://leetcode-cn.com/problems/basic-calculator-ii/
// 基本计算器 II
// 解题思路：用数组模拟栈
func calculate(s string) int {
	var tempSlice []int //临时数组，记录所有待相加元素
	var num int
	preSign := '+'
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit { //记录数字
			num = num*10 + int(ch-'0')
		}
		if (!isDigit && ch != ' ') || i == len(s)-1 { //ch为符号 || 遍历到最后一个字符
			switch preSign {
			case '+': //直接加入数组
				tempSlice = append(tempSlice, num)
			case '-': //将负数加入数组
				tempSlice = append(tempSlice, -num)
			case '*': // 将数组最后一个拿出来*num，将结果压入数组
				tempSlice[len(tempSlice)-1] *= num
			case '/': // 将数组最后一个拿出来/num，将结果压入数组
				tempSlice[len(tempSlice)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	var sum int
	for i := range tempSlice { //将临时数组所有元素相加
		sum += tempSlice[i]
	}
	return sum
}
