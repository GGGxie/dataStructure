package svc

import (
	"shorturl/api/internal/config"
	"shorturl/rpc/transform/transformclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformclient.Transform // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformclient.NewTransform(zrpc.MustNewClient(c.Transform)),
	}
}
