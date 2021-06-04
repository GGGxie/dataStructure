package main

import (
	"fmt"
)

func main() {
	fmt.Println(printBin(0.7))
	// fmt.Println(1.423%10)
	// fmt.Printf("%016b\n", bits.InsertBits(0b11111, 0b110, 1, 4))
}
func printBin(num float64) string {
	var res []byte
	res = append(res, '0')
	res = append(res, '.')

	for num != 0 {
		num *= 2
		fmt.Println(num)
		if num >= 1 {
			res = append(res, '1')
			num -= 1
		} else {
			res = append(res, '0')
		}
	}
	fmt.Println(string(res))
	if len(res) >= 32 {
		return "ERROR"
	}
	return string(res)
}
