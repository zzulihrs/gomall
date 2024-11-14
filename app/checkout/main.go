package main

import (
	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	consul "github.com/kitex-contrib/registry-consul"
	"net"
)

func main() {

	_ = godotenv.Load()
	rpc.InitClient()

	//dal.Init()
	rpc.InitClient()

	opts := kitexInit()

	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}), server.WithRegistry(r))
	return
}
