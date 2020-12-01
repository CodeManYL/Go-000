package logic

import (
	"errors"
	"fmt"
	"github.com/CodeManYL/Go-000/Week02/dao"
	"github.com/gin-gonic/gin"
)

func UserHandle(c *gin.Context){
	userid := c.Query("userid")

	user,err := dao.GetUserInfo()
	if errors.Is(err,dao.ErrorQueryNotExist) {
		c.JSON(200, gin.H{
			"message": dao.ErrorQueryNotExist.Error(),
		})
		return
	}

	if err != nil {
		//记录错误日志
		fmt.Println(fmt.Sprintf("userid:%s;err:%+v",userid,err))
	}

	c.JSON(200, gin.H{
		"message": user,
	})
}