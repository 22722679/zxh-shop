package frontend

import (
	"context"
	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/frontend/auth"
	common "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/frontend/common"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Login(ctx context.Context, req *auth.LoginReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Register(ctx context.Context, req *auth.RegisterReq, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Logout(ctx context.Context, req *common.Empty, callOptions ...callopt.Option) (resp *common.Empty, err error) {
	resp, err = defaultClient.Logout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
