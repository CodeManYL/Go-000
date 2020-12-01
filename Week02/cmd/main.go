package main

import (
	"github.com/CodeManYL/Go-000/Week02/logic"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/test", logic.UserHandle)
	r.Run()
}

