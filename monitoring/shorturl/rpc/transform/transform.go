package main

import (
	"flag"
	"fmt"

	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/internal/config"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/internal/server"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/internal/svc"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/transform"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTransformerServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformerServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
