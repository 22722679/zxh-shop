package frontend

import (
	"context"
	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/frontend/auth"
	common "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/frontend/common"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/frontend/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	Login(ctx context.Context, Req *auth.LoginReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	Register(ctx context.Context, Req *auth.RegisterReq, callOptions ...callopt.Option) (r *common.Empty, err error)
	Logout(ctx context.Context, Req *common.Empty, callOptions ...callopt.Option) (r *common.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Login(ctx context.Context, Req *auth.LoginReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) Register(ctx context.Context, Req *auth.RegisterReq, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Logout(ctx context.Context, Req *common.Empty, callOptions ...callopt.Option) (r *common.Empty, err error) {
	return c.kitexClient.Logout(ctx, Req, callOptions...)
}
