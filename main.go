package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

// https://leetcode-cn.com/problems/gas-station/
// 加油站
// 模拟题
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(cost)
	tempList := make([]int, length)    //标记从i节点开始，到最后因为缺少多少汽油而停止
	for i := length - 1; i >= 0; i-- { //从最后一个节点开始遍历，然后往前去遍历每一个节点开始的情况（因为这样子后面节点的tempList都有值了）
		if gas[i] < cost[i] {
			continue
		}
		tempGas := 0 //初始汽油值
		count := 0
		for j := i; ; j = ((j + 1) % length) { //从i节点开始往前遍历
			count++
			tempGas += gas[j]
			tempGas -= cost[j]
			if tempGas < 0 { //如果从i点开始，到了j点，汽油不够了，就记录到tempList中
				tempList[i] = tempGas
				break
			}
			if tempGas < tempList[j] { //如果从i点开始，到了j点，汽油还充足，但是少于从j开始到最后缺少的汽油，则没必要继续遍历
				tempList[i] = tempGas - tempList[j]
				break
			}
			if count == length {
				fmt.Println(tempList)
				return i
			}
		}

	}
	return -1
}
