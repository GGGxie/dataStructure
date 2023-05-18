package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(FirstMiddleware())
	router.Use(SecondMiddleware(), ThirdMiddleware()) // (1)
	router.GET("test", H)
	router.Run()
}

func H(c *gin.Context) {
	c.String(http.StatusOK, `{
		"title": {
		  "text": "已处理告警数"
		},
		"tooltip": {},
		"xAxis": {
		  "data": "${name_list}"
		},
		"yAxis": {},
		"series": [
		  {
			"realtimeSort": true,
			"name": "销量",
			"type": "bar",
			"data": "${value_list}"
		  }
		]
	  }`)
}

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			c.Abort()
		}()
		fmt.Println("")
		c.Next()
		return
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of second middleware")
	}
}

func ThirdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of third middleware")
	}
}
