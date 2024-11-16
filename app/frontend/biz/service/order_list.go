package service

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/types"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"time"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint32(userId),
	})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)

		for _, item := range v.Items {
			total += item.Cost
			i := item.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
				Id: i.ProductId,
			})
			if err != nil {
				return nil, err
			}

			if productResp == nil || productResp.Product == nil {
				continue
			}

			items = append(items, types.OrderItem{
				ProductName: productResp.Product.Name,
				Picture:     productResp.Product.Picture,
				Qty:         i.Quantity,
				Cost:        item.Cost,
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
