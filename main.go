package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcd"
	"go-micro.dev/v4/registry"
)

func main() {
	// 创建 etcd 注册中心实例
	etcdReg := etcd.NewRegistry(
		registry.Addrs("localhost:2379"),
	)

	// 创建一个新的服务客户端
	client := micro.NewService(
		micro.Registry(etcdReg),
	)

	// 创建服务客户端
	myServiceClient := myService.NewMyServiceClient("my.service.name", client.Client())

	// 调用服务
	resp, err := myServiceClient.MyMethod(context.TODO(), &myService.Request{})
	if err != nil {
		fmt.Println("调用服务失败:", err)
		return
	}

	fmt.Println("服务调用结果:", resp)
}
