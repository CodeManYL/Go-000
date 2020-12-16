package main

import (

	"fmt"
	user "github.com/CodeManYL/Go-000/Week03/app/service/users/api"
	"github.com/CodeManYL/Go-000/Week03/app/service/users/configs"
	"github.com/CodeManYL/Go-000/Week03/app/service/users/internal/biz"
	"github.com/CodeManYL/Go-000/Week03/app/service/users/internal/data"
	"github.com/CodeManYL/Go-000/Week03/app/service/users/internal/service"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/transport/grpc"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

func serviceRegister(conf *configs.UserRpcConf,s micro.Service){
	userRepo,err := data.NewUserData(conf.Db.Name,conf.Db.Address)
	if err != nil {
		panic(err)
	}
	userBiz := biz.NewUserBiz(userRepo)
	userService := service.NewUserService(userBiz)

	if err := user.RegisterUserHandler(s.Server(),userService);err != nil {
		panic(err)
	}
}

func main() {
	//初始化配置文件
	conf,err := configs.InitConfig()
	if err != nil {
		panic(fmt.Sprintf("init config failed err :%v",err))
	}

	etcd := etcdv3.NewRegistry(
		registry.Addrs(conf.Etcd.Address),
	)

	// Create a new service. Optionally include some options here.
	s := micro.NewService(
		micro.Name(conf.ServerName),
		micro.Transport(grpc.NewTransport()),
		micro.Registry(etcd),
		micro.Address(conf.ServerAddr),
		micro.Flags(&cli.StringFlag{
			Name:  "e",
			Value: "./config/config_rpc.json",
			Usage: "please use xxx -f config_rpc.json",
		}),
	)
	s.Init()

	serviceRegister(conf,s)

	if err := s.Run(); err != nil {
		panic(err)
	}

}