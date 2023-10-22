package gateway

import (
	"github.com/GGGxie/dataStructure/microservice/micro/service/gateway/api/service_a"
	"github.com/GGGxie/dataStructure/microservice/micro/service/gateway/api/service_b"
	"github.com/GGGxie/dataStructure/microservice/micro/service/gateway/router"
)

func NewGateway() *Gateway {
	return &Gateway{}
}

type Gateway struct {
}

func (g *Gateway) Run(existChan chan bool) error {
	InitClient()

	var err error
	routerOb := router.GetRouter()
	// set middleware
	// routerOb.Use()
	Register(routerOb)
	routerOb.Setup()
	err = routerOb.ListenAndServe(":8080")
	return err
}

// 注册服务
func Register(router *router.Router) {
	// 注册服务 A
	router.Register(service_a.NewSARouter("", ""))
	// 注册服务 B
	router.Register(service_b.NewSBRouter("", ""))
}

// 注册 rpc 客户端
func InitClient() {
	service_a.InitClient()
	service_b.InitClient()
}
