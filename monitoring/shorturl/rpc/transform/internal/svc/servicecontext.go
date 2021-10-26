package svc

import (
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/internal/config"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	c     config.Config
	Model model.ShorturlModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
