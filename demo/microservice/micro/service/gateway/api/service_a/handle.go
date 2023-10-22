package service_a

import (
	"context"
	"net/http"

	p "github.com/GGGxie/dataStructure/microservice/micro/proto/service_a"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Query("Name")
	resp, err := saClient.Hello(context.Background(), &p.Request{
		Name: name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, resp)
}
