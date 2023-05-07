package router

import (
	"go-chat/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	middleware.InitMiddleware(r)

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// // swagger
	// docs.SwaggerInfo.BasePath = ""
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// // 首页
	// r.GET("/", service.GetIndex)

	// // 用户模块
	// r.POST("/user", service.CreateUser)
	// r.DELETE("/user/:id", service.DeleteUser)
	// r.GET("/user/list", service.GetUserList)

	// // 登录
	// r.POST("login", service.Login)

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	return r
}
