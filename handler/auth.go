package handler

import (
	"fmt"
	"go-chat/models"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"pssword"`
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.UserBasic); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v.Name,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.UserBasic{
		Name: claims["Name"].(string),
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginUser LoginUser
	if err := c.ShouldBind(&loginUser); err != nil {
		return "", jwt.ErrMissingLoginValues
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		fmt.Printf("v: %v\n", v)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
