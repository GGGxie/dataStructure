package router

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var router *Router
var once sync.Once

type Rs interface {
	GetBasePath() string
	SetRouter(groupBaseUrl string, ge *gin.Engine)
}

// 单例模式
func GetRouter() *Router {
	once.Do(func() {
		router = &Router{
			RList:   make(map[string][]Rs),
			GEngine: gin.Default(),
		}
	})
	return router
}

type Router struct {
	RList   map[string][]Rs
	GEngine *gin.Engine
}

// 使用中间件
func (r *Router) Use(middleware ...gin.HandlerFunc) {
	r.GEngine.Use(middleware...)
}

// 注册路由
func (r *Router) Register(rs Rs) {
	bPath := rs.GetBasePath()
	r.RList[bPath] = append(r.RList[bPath], rs)
}

// 设置路由
func (r *Router) Setup() {
	for url, rl := range r.RList {
		for _, rtmp := range rl {
			rtmp.SetRouter(url, r.GEngine)
		}
	}
}

func (r *Router) ListenAndServe(addr string) error {
	return r.GEngine.Run(addr)
}
