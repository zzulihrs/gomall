package service

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/common/utils"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	rpccheckout "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Email:     req.Email,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil
}
