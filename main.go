package main

import "fmt"

func main() {
	fmt.Println(reverseBits(2147482622))
}
func reverseBits(num int) int {
	var pre, max, current int
	var flag bool
	for i := 0; i < 32; i++ { //遍历num
		a := num & 1
		num = num >> 1
		if a == 1 {
			current++
			pre++
		} else if a == 0 {
			if !flag {
				current++
				flag = true
			} else {
				current = pre + 1
				if current > max {
					max = current
				}
			}
			pre = 0
		}
		if current > max {
			max = current
		}
	}
	return max
}
