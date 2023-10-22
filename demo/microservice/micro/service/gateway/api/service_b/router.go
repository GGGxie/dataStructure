package service_b

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	moduleName    = "service_b"
	routerVersion = "v1"
)

func NewSBRouter(name, version string) *SBRouter {
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
	return &SBRouter{
		ModuleName:    mName,
		RouterVersion: rVersion,
	}
}

type SBRouter struct {
	ModuleName    string
	RouterVersion string
}

func (sb *SBRouter) GetBasePath() string {
	return fmt.Sprintf("%s/%s", sb.ModuleName, sb.RouterVersion)
}

func (sa *SBRouter) SetRouter(url string, e *gin.Engine) {
	group := e.Group(url)
	group.GET("/hello/", Hello)
}
