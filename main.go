package main

import (
	"fmt"
)

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
	sql := fmt.Sprintf("object like '%%name: %s%%' OR object like '%%name: %s%%'", "sdf", "twe")
	fmt.Println(sql)
}
