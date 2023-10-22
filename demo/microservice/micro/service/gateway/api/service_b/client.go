package service_b

import (
	"github.com/GGGxie/dataStructure/microservice/micro/pkg/consts"
	p "github.com/GGGxie/dataStructure/microservice/micro/proto/service_b"
	"github.com/go-micro/plugins/v4/wrapper/breaker/hystrix"
	"go-micro.dev/v4"
)

var (
	sbClient p.ServiceBService
)

func InitClient() {
	// create a new service
	hystrix.ConfigureDefault(
		hystrix.CommandConfig{
			//超时 5000ms
			Timeout: 5000,
			// 最大请求并发数 100
			MaxConcurrentRequests: 100,
			// 错误百分比阈值为 25%
			ErrorPercentThreshold: 25,
		},
	)
	service := micro.NewService(
		micro.Name("b.client"),
		// 用于包装对外发出的请求，即客户端包装
		micro.WrapClient(hystrix.NewClientWrapper()),
	)

	// parse command line flags
	service.Init(
		micro.Address("39.108.148.65:40659"),
	)
	sbClient = p.NewServiceBService(consts.ServiceB, service.Client())
}
