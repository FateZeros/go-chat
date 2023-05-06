package router

import (
	"go-chat/docs"
	"go-chat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 首页
	r.GET("/", service.GetIndex)

	// 用户模块
	r.POST("/user", service.CreateUser)
	r.DELETE("/user/:id", service.DeleteUser)
	r.GET("/user/list", service.GetUserList)

	return r
}
