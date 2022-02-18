package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"ihomeMs/service/getImg/handler"

	"github.com/micro/go-micro/registry/consul"
	"ihomeMs/service/getImg/model"
	getImg "ihomeMs/service/getImg/proto/getImg"
)

func main() {
	model.InitRedis()

	//初始化consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"47.94.195.58:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Address("172.16.161.25:12332"), //云服务器使用自己的consul只需要绑定本地的本机ip
		//micro.Address("192.168.31.12:12332"), //windows 本机
		micro.Name("go.micro.srv.getImg"),
		micro.Registry(consulReg), // 添加注册
		micro.Version("v1.7.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getImg.RegisterGetImgHandler(service.Server(), new(handler.GetImg))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
