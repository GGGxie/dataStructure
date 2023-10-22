package service_b

import (
	"context"
	"net/http"

	p "github.com/GGGxie/dataStructure/microservice/micro/proto/service_b"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Query("name")
	resp, err := sbClient.Hello(context.Background(), &p.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
