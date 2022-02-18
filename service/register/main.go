package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"ihomeMs/service/register/handler"
	"ihomeMs/service/register/model"
	register "ihomeMs/service/register/proto/register"
)

func main() {
	//初始化consul
	consulReg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"47.94.195.58:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Address("172.16.161.25:12334"), //云服务器使用自己的consul只需要绑定本地的本机ip
		//micro.Address("192.168.31.12:12332"), //windows 本机
		micro.Name("go.micro.srv.register"),
		micro.Registry(consulReg), // 添加注册
		micro.Version("v1.7.0"),
	)

	// Initialise service
	service.Init()
	model.InitRedis()
	model.InitDb()

	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
