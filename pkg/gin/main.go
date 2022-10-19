package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.Run(":8080")
}

func testHandler(c *gin.Context) {
	c.Redirect(302, `dingtalk://dingtalkclient/page/link?url=https%3A%2F%2Fdevops-stage.onething.net%2Fo%2Fnew-itsm2%2F%23%2Fticket%2Fdetail%3Fid%3D2483%26from%3Dcreated&pc_slide=true`)
}
