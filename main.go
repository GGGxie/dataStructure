package main

import "fmt"

func main() {
	fmt.Println(reverseBits(0b110101111))
}

//翻转数位,你可以将一个数位从0变为1,找出你能够获得的最长的一串1的长度。TODO:改成动态规划
func reverseBits(num int) int {
	//pre记录从上一个0开始，1的总数
	//max记录最大的1的总数
	//current记录用0连接的两段的1的总数
	var pre, max, current int
	for i := 0; i < 32; i++ { //遍历num
		a := num & 1 //取出个位
		num = num >> 1
		if a == 1 {
			current++
			pre++
		} else if a == 0 {
			current = pre + 1
			pre = 0
		}
		if current > max {
			max = current
		}
	}
	return max
}
