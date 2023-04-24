package router

import (
	"go-chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//首页
	r.GET("/", service.GetIndex)

	return r
}
