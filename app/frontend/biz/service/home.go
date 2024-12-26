package service

import (
	"context"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	// resp := make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-2", "Price": 120, "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-3", "Price": 50, "Picture": "/static/image/t-shirt-2.jpg"},
	// 	{"Name": "T-shirt-4", "Price": 70, "Picture": "/static/image/t-shirt-1.jpg"},
	// 	{"Name": "T-shirt-5", "Price": 80, "Picture": "/static/image/t-shirt-2.jpg"},
	// 	{"Name": "T-shirt-6", "Price": 100, "Picture": "/static/image/t-shirt.png"},
	// }
	// resp["title"] = "Hot Sales"
	// resp["items"] = items
	//return resp, nil
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Hot sale",
		"items": products.Products,
	}, nil
}
