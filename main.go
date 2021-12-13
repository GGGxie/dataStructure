package main

import "fmt"

var ret int

func main() {
	str := "123"
	str = "1234"
	str = "12345"  //12 34 5, 123 45
	str = "123456" //12 34 56,12 34 5 6,123 45 6
	str = "123456789"
	Count(str)
	fmt.Println(ret)
}
