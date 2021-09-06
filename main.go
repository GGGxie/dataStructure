package main

import (
	"fmt"
)

func f(i int) error {
	return nil
}

func f2(f func(i int) error) {

}

func main() {
	a := 0x10000011
	z := int8(a)
	c := uint8(a)
	fmt.Printf("%08b\n", z)
	fmt.Printf("%08b\n", c)
}
