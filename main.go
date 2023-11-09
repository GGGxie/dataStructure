package main

import "fmt"

func main() {
	mapp := make(map[int]int)
	mapp[1] = 1
	fmt.Println(len(mapp))
}
