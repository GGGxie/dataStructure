package service_a

import (
	"context"
	"fmt"

	"github.com/GGGxie/dataStructure/microservice/micro/pkg/consts"
	s "github.com/GGGxie/dataStructure/microservice/micro/service/service_a/service"
	"go-micro.dev/v4"
)

func NewService() *Service {
	return &Service{}
}

type Service struct {
}

func (svc *Service) Run(existChan chan bool) error {
	ctx, cancel := context.WithCancel(context.Background())
	service := micro.NewService(
		micro.Name(consts.ServiceA),
		micro.Context(ctx),
		micro.Handle(new(s.ServiceA)),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()
	go func() {
		<-existChan
		cancel()
		fmt.Println("Service A exit!")
	}()
	return nil
}
