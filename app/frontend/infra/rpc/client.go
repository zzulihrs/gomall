package rpc

import (
	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once // 保证只初始化一次
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
	})
}

func initUserClient() {

	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}
