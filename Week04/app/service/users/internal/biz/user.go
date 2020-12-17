package biz

import (
	"fmt"
)

type UserMod interface {
	 GetUserByUsernameAndPassword(username,password string) (*UserInfo,error)
}

type UserBiz struct {
	userMod UserMod
}

func NewUserBiz(userRepo UserMod) *UserBiz {
	return &UserBiz{userRepo}
}

func (u *UserBiz) Login(username,password string) (bool,error){
	res,err := u.userMod.GetUserByUsernameAndPassword(username,password)
	if err != nil {
		return false,err
	}
	fmt.Println(res)
	return true,nil
}