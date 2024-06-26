package logic

import (
	"context"

	"github.com/GGGxie/dataStructure/monitoring/shorturl/api/internal/svc"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/api/internal/types"
	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/transformer"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandLogic {
	return ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (types.ExpandResp, error) {
	// 手动代码开始
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transformer.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return types.ExpandResp{}, err
	}

	return types.ExpandResp{
		Url: resp.Url,
	}, nil
	// 手动代码结束
}
