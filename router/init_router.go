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

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	return r
}
