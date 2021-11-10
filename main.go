package main

import (
	"fmt"
	"math"
)

func main() {
	pre := [][]int{{0, 1}, {0, 0}, {0, 4}, {0, -2}, {0, -1}, {0, 3}, {0, -4}}
	// pre := {}{}int{{0, 1}, {3, 1}, {1, 3}, {3, 2}}
	// pre := {}{}int{{1, 0}}
	fmt.Println(math.Inf(1))
	fmt.Println(maxPoints(pre))
}
