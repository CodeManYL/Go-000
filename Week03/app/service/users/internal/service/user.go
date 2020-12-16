package service

import (
	"context"
	user "github.com/CodeManYL/Go-000/Week03/app/service/users/api"
	"github.com/CodeManYL/Go-000/Week03/app/service/users/internal/biz"

	"fmt"
	"github.com/micro/go-micro/v2/errors"
)

type User struct{
	userBiz *biz.UserBiz
}

func NewUserService(userBiz *biz.UserBiz) *User {
	return &User{userBiz}
}

func (u *User) Login(ctx context.Context, req *user.Request, rsp *user.Response) error {
	//参数校验
	//.......
	//登陆逻辑
	ok, err := u.userBiz.Login(req.Username,req.Password)
	if err != nil {
		fmt.Println(err)
		return errors.NotFound("user.server.v1","not found")
	}
	if !ok {
		return errors.NotFound("user.server.v1","not found")
	}
	//token handle
	//.......

	rsp.Ok = true
	return nil
}
