package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var items []map[string]any
	var total float64
	if cartResp != nil && cartResp.Cart != nil {

		for _, item := range cartResp.Cart.Items {
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
				Id: item.ProductId,
			})
			if err != nil {
				return nil, err
			}
			p := productResp.Product
			items = append(items, map[string]any{
				"Name":        p.Name,
				"Description": p.Description,
				"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
				"Picture":     p.Picture,
				"Qty":         strconv.Itoa(int(item.Quantity)),
			})
			total += float64(p.Price) * float64(item.Quantity)
		}
	} else {
		// 处理 cartResp 或 cartResp.Cart 为 nil 的情况
		fmt.Println("cartResp or cartResp.Cart is nil")
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
