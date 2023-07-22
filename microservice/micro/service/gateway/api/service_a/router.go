package service_a

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	moduleName    = "service_a"
	routerVersion = "v1"
)

func NewSARouter(name, version string) *SARouter {
	var mName, rVersion string
	if name != "" {
		mName = name
	} else {
		mName = moduleName
	}
	if version != "" {
		rVersion = version
	} else {
		rVersion = routerVersion
	}
	return &SARouter{
		ModuleName:    mName,
		RouterVersion: rVersion,
	}
}

type SARouter struct {
	ModuleName    string
	RouterVersion string
}

func (sa *SARouter) GetBasePath() string {
	return fmt.Sprintf("%s/%s", sa.ModuleName, sa.RouterVersion)
}

func (sa *SARouter) SetRouter(url string, e *gin.Engine) {
	group := e.Group(url)
	group.GET("/hello/", Hello)
}
