// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"example/engineering-layout/app/service/users/configs"
	"example/engineering-layout/app/service/users/internal/biz"
	"example/engineering-layout/app/service/users/internal/data"
	"example/engineering-layout/app/service/users/internal/service"
)

// Injectors from wire.go:

func InitializeService(cof *configs.UserRpcConf) (*service.User, error) {
	userMod, err := data.NewUserData(cof)
	if err != nil {
		return nil, err
	}
	userBiz := biz.NewUserBiz(userMod)
	user := service.NewUserService(userBiz)
	return user, nil
}
