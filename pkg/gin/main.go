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
	router.GET("auth", Auth)
	router.Run()
}

func Auth(c *gin.Context) {
	forwardURL := c.Request.FormValue("forward_url")
	code := c.Request.FormValue("authCode")
	state := c.Request.FormValue("state")
	redirectURL := fmt.Sprintf("%s?code=%s&state=%s", forwardURL, code, state)
	fmt.Println(redirectURL)
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
