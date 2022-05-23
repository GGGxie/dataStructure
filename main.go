package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResp struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Result  bool        `json:"result"`
}

var db = make(map[string]string)

func setupRouter() *gin.Engine {

	r := gin.Default()

	// Ping test
	r.POST("/test", func(c *gin.Context) {
		commonResp := &CommonResp{
			Result:  true,
			Message: "",
			Data: struct {
				Msg string `json:"msg"`
			}{"tess"},
			Code: 200,
		}
		c.JSON(http.StatusOK, commonResp)
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
