package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"ihomeMs/service/getArea/handler"
	"ihomeMs/service/getArea/model"
	getArea "ihomeMs/service/getArea/proto/getArea"
)

func main() {
	model.InitDb()
	model.InitRedis()
	//初始化consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"47.94.195.58:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Address("172.16.161.25:12331"), //云服务器使用自己的consul只需要绑定本地的本机ip
		//micro.Address("192.168.31.12:12332"), //windows 本机
		micro.Name("go.micro.srv.getArea"),
		micro.Registry(consulReg), // 添加注册
		micro.Version("v1.7.0"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
