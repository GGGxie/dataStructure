package service_a

import (
	"github.com/GGGxie/dataStructure/microservice/micro/pkg/consts"
	p "github.com/GGGxie/dataStructure/microservice/micro/proto/service_a"
	"go-micro.dev/v4"
)

var (
	saClient p.ServiceAService
)

func InitClient() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	saClient = p.NewServiceAService(consts.ServiceA, service.Client())
}
