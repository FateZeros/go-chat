package service

import (
	"fmt"
	"go-chat/models"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"pssword"`
}

// Login
// @Summary 登录
// @Tags 登录
// @Params
// @Success 200 {string} json{"code", "message"}
// @Router /login [post]
func Login(c *gin.Context) {
	var loginUser LoginUser
	c.BindJSON(&loginUser)

	fmt.Printf("loginUser: %v\n", loginUser)

	user := models.FindUserByName(loginUser.Name)

	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户不存在",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "登录成功",
		"data":    "",
	})

}
