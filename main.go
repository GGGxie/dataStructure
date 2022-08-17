package main

import (
	"flag"
	"fmt"

	"github.com/GGGxie/dataStructure/A"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	A.T()
	a := [1]int{3}
	Ze(&a)
	fmt.Printf("%T %[1]v\n", a) // "time.Duration 0"
}

func Ze(z *[1]int) {
	z[0] = 1
}
