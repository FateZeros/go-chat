package router

import (
	"go-chat/docs"
	"go-chat/handler"

	systemRouter "go-chat/router/system"

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

	// 需要认证的接口
	sysAuthRouterInit(g, authMiddleware)

	return g
}

func sysSwaggerRouter(g *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = ""
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func SysBaseRouter(g *gin.RouterGroup) {
	g.GET("/ping", handler.Ping)
}

func sysAuthRouterInit(g *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	g.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	g.GET("/refresh_token", authMiddleware.RefreshHandler)

	v1 := g.Group("/api/v1")

	// 用户
	systemRouter.RegisterUserRouter(v1, authMiddleware)
}
