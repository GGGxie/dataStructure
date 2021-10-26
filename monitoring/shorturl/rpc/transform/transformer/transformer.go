// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package transformer

import (
	"context"

	"github.com/GGGxie/dataStructure/monitoring/shorturl/rpc/transform/transform"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	ExpandReq   = transform.ExpandReq
	ExpandResp  = transform.ExpandResp
	ShortenReq  = transform.ShortenReq
	ShortenResp = transform.ShortenResp

	Transformer interface {
		Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error)
		Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error)
	}

	defaultTransformer struct {
		cli zrpc.Client
	}
)

func NewTransformer(cli zrpc.Client) Transformer {
	return &defaultTransformer{
		cli: cli,
	}
}

func (m *defaultTransformer) Expand(ctx context.Context, in *ExpandReq) (*ExpandResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Expand(ctx, in)
}

func (m *defaultTransformer) Shorten(ctx context.Context, in *ShortenReq) (*ShortenResp, error) {
	client := transform.NewTransformerClient(m.cli.Conn())
	return client.Shorten(ctx, in)
}
