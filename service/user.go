package service

import (
	"fmt"
	"go-chat/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Accept application/json
// @Param data body models.UserBasic true "data"
// @Success 200 {string} json"{"code": 200,"message": "新增成功"}"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	userData := models.UserBasic{}
	c.BindJSON(&userData)

	user := models.UserBasic{}

	user.Name = userData.Name
	password := userData.PassWord

	fmt.Printf("user.Name: %v\n", user.Name)

	data := models.FindUserByName(user.Name)

	if user.Name == "" || password == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或者密码不能为空",
			"data":    user,
		})
		return
	}

	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名已注册",
			"data":    user,
		})
		return
	}

	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功！",
		"data":    user,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param userId path string true "id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/{userId} [delete]
func DeleteUser(c *gin.Context) {

}
