package svc

import (
	"shorturl/rpc/transform/internal/config"

	"shorturl/rpc/transform/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // manual code
	}
}
