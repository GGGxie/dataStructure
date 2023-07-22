package service

import (
	"context"
	"fmt"

	pb "github.com/GGGxie/dataStructure/microservice/micro/proto/service_a"
)

type ServiceA struct{}

func (a *ServiceA) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = fmt.Sprintf("A:%s", req.Name)
	return nil
}
