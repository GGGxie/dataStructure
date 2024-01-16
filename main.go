package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func A[T, C int | float64]() T {
	return 1
}
func main() {
	app := gin.Default()
	// 指明html加载文件目录
	app.LoadHTMLGlob("./html/*")
	app.Handle("GET", "/", func(context *gin.Context) {
		// 返回HTML文件，响应状态码200，html文件名为index.html，模板参数为nil
		context.HTML(http.StatusOK, "pay.html", nil)
	})
	app.Run()
}
