package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	s int
)

func init() {
	flag.IntVar(&s, "s", 0, "sleep time")
}
func main() {
	flag.Parse()
	if s != 0 {
		fmt.Printf("Service will sleep %d second\n", s)
		time.Sleep(time.Second * time.Duration(s))
	}

	router := gin.Default()
	router.GET("health", Health)
	router.GET("message/:n", H)
	router.GET("cancel")
	router.Run("0.0.0.0:8080")
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "Service is healthy!")
}
func Cancel(c *gin.Context) {
	c.JSON(http.StatusOK, "Service will be done!")
	os.Exit(0)
}
func H(c *gin.Context) {
	n := c.Query("n")
	nint, _ := strconv.ParseInt(n, 10, 64)
	ret := new(int)
	hannuota(nint, ret)
	c.JSON(http.StatusOK, map[interface{}]interface{}{"call": n, "result": *ret})
}

// 汉诺塔问题
func hannuota(n int64, num *int) {
	//就剩一个盘子的，直接将这个盘子从起始盘移到目标盘
	if n == 1 {
		// fmt.Println("将第", n, "个盘子从", from, "移到", to)
	} else {
		//将上面的n-1个盘子从起始盘移到中转盘
		hannuota(n-1, num)
		//将最下面的第n个盘子从起始盘移到目标盘
		// fmt.Println("将第", n, "个盘子从", from, "移到", to)
		//再将中转盘上面的n-1个盘子移到目标盘
		hannuota(n-1, num)
	}
	*num++
}
