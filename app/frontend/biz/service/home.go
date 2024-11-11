package service

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {

	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	var cartNum int
	//items := []struct {
	//	Id      int
	//	Name    string
	//	Price   float64
	//	Picture string
	//}{
	//	{Id: 1, Name: "商品A", Price: 99.99, Picture: "/static/image/t-shirt-1.jpeg"},
	//	{Id: 2, Name: "商品B", Price: 199.99, Picture: "/static/image/t-shirt-2.jpeg"},
	//	{Id: 3, Name: "商品C", Price: 299.99, Picture: "/static/image/t-shirt-bag.jpeg"},
	//}
	return utils.H{
		"title":    "Hot sale",
		"cart_num": cartNum,
		"items":    products.Products,
	}, nil
}
