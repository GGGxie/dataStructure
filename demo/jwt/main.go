package main

import (
	"fmt"
	"time"

	pkg "github.com/GGGxie/dataStructure/pkg/jwt"
)

func main() {
	token, err := pkg.GenerateToken("root", "pawd")
	fmt.Println(token, err)
	time.Sleep(2 * time.Second)
	claims, err := pkg.ParseToken(token)
	fmt.Println(claims, err)
}
