package main

import (
	"flag"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/server/grpc/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"meshop-product-service/application/handler"
	sysConfig "meshop-product-service/config"
	"meshop-product-service/datasource"
	"meshop-product-service/middleware"
)

func InitServer() {
	//解析命令运行参数
	flag.Parse()
	//从配置中心获取业务配置
	sysConfig.InitSysConfig()
	//连接数据库
	datasource.InitDatabase()
}

func main() {
	InitServer()

	conf := sysConfig.Get()
	log.Info("Version:", conf.Service.Version)

	//注册
	consulRegistry := consul.NewRegistry(registry.Addrs(conf.Consul.Addrs))
	opts := []micro.Option{
		micro.Server(grpc.NewServer()),
		micro.Address(conf.Service.ListenHost()),
		micro.Name(conf.Service.ServiceName),
		micro.Version(conf.Service.Version),
		micro.Registry(consulRegistry),
		micro.WrapHandler(middleware.RecoverWrapper),
		micro.WrapClient(middleware.NewClientWrapper()),
	}

	service := micro.NewService(opts...)
	service.Init()

	//注册grpc handler
	handler.RegisterProduct(service)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
