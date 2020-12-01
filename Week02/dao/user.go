package dao

import (
	"database/sql"
	"errors"
	errorsx "github.com/pkg/errors"
)

//pgk包里的New创建的对战会记录堆栈,故不用
//这里就定义在这里不单独弄包了
var  ErrorQueryNotExist = errors.New("查询不存在")

type User struct {
	Name string
	Age  int
}

func GetUserInfo()(user *User,err error){
	sqlStr := "select * from user where id = 1"
	//对数据库的操作省略
	err = sql.ErrNoRows
	//空查询不记录，用户登陆等尝试性行为，没有保存堆栈记录日志的必要
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,ErrorQueryNotExist
		}
		return nil,errorsx.Wrap(err,sqlStr)
	}

	return
}
