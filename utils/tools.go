package utils

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func GetMicroClient() client.Client {
	//从consul中获取服务
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"47.94.195.58:8500",
		}
	})
	microService := micro.NewService(
		micro.Registry(consulReg),
	)
	return microService.Client()
}
