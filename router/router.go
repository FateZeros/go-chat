package router

import (
	"go-chat/docs"
	"go-chat/handler"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	SysBaseRouter(g)

	// swagger router
	sysSwaggerRouter(g)

	return g
}

func sysSwaggerRouter(g *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = ""
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func SysBaseRouter(g *gin.RouterGroup) {
	g.GET("/ping", handler.Ping)
}
