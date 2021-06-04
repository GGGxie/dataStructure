package main

import (
	"fmt"
)

func main() {
	var z int
	z = 12
	b := 4
	fmt.Printf("%016b\n", 1<<2)
	fmt.Printf("%016b\n", ^(1 << 2))
	fmt.Printf("%016b\n", z)
	fmt.Printf("%016b\n", z|(^(1 << 2)))
	fmt.Printf("%016b\n", z&(^b))
	// fmt.Printf("%016b\n", bits.InsertBits(0b11111, 0b110, 1, 4))
}
