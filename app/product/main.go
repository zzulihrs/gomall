package main

import (
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/app/product/conf"
	"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"github.com/cloudwego/biz-demo/gomall/common/serversuite"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	"net"
)

var (
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddr       = conf.GetConf().Registry.RegistryAddress[0]
	MetricsPort        = conf.GetConf().Kitex.MetricsPort
)

func main() {
	_ = godotenv.Load()
	mtl.InitMetric(CurrentServiceName, MetricsPort, RegistryAddr)

	dal.Init()
	opts := kitexInit()

	svr := productcatalogservice.NewServer(new(ProductCatalogServiceImpl), opts...)

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

	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	return
}
