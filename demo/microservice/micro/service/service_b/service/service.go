package service

import (
	"context"
	"fmt"

	pb "github.com/GGGxie/dataStructure/microservice/micro/proto/service_b"
)

type ServiceB struct{}

func (b *ServiceB) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = fmt.Sprintf("B:%s", req.Name)
	return nil
}
