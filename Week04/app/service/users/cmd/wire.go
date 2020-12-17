// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"example/engineering-layout/app/service/users/configs"
	"example/engineering-layout/app/service/users/internal/biz"
	"example/engineering-layout/app/service/users/internal/data"
	"example/engineering-layout/app/service/users/internal/service"
	"github.com/google/wire"
)

func InitializeService(cof *configs.UserRpcConf) (*service.User,error) {
	wire.Build(service.NewUserService,biz.NewUserBiz,data.MockUserRepoSet)
	return &service.User{},nil
}