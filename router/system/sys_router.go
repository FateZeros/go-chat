package system

import (
	apisSystem "go-chat/apis/system"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	user := v1.Group("/user").Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("", apisSystem.CreateUser)
		user.DELETE("/:userId", apisSystem.DeleteUser)
		user.GET("/list", apisSystem.GetUserList)
	}
}
