package service_b

import (
	"context"
	"fmt"

	"github.com/GGGxie/dataStructure/microservice/micro/pkg/consts"
	pb "github.com/GGGxie/dataStructure/microservice/micro/proto/service_b"
	s "github.com/GGGxie/dataStructure/microservice/micro/service/service_b/service"
	limiter "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	opentracingplugins "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (svc *Service) Run(existChan chan bool) error {
	initJaeger()
	ctx, cancel := context.WithCancel(context.Background())
	service := micro.NewService(
		micro.Name(consts.ServiceB),
		micro.Context(ctx),
		// micro.Handle(new(s.ServiceB)),
		// 添加限流
		micro.WrapHandler(
			// 用于包装外界发来的请求，即服务端包装
			// QPS上限为 100
			limiter.NewHandlerWrapper(100),
		),
		// 添加链路追踪
		micro.WrapHandler(opentracingplugins.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	pb.RegisterServiceBHandler(service.Server(), new(s.ServiceB))

	// initialise flags
	service.Init()

	// start the service
	go func() {
		service.Run()
	}()
	go func() {
		<-existChan
		cancel()
		fmt.Println("Service B exit!")
		existChan <- true
	}()
	return nil
}

func initJaeger() {
	zipkin.NewJaegerTracer(consts.ServiceB, "39.108.148.65:16686")
}
