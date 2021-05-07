package main

import "fmt"

func main() {
	a := []int{1, 3}
	fmt.Println(a[0:])
	fmt.Println(a[1:])
}
func te(t interface{}) {
	fmt.Println(t == nil)
}
