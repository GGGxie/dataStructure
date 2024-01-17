package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("test", H)
	router.Run(":30000")
}

func H(c *gin.Context) {
	c.String(http.StatusOK, `Hello World`)
}
