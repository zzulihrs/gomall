package rpc

import (
	"github.com/cloudwego/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client

	once sync.Once // 保证只初始化一次
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {

	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))

	frontendUtils.MustHandleError(err)
}
