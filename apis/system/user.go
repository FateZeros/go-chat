package system

import (
	"fmt"
	"go-chat/models"
	"strconv"
	"time"

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

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Accept application/json
// @Param userData body models.UserBasic true "userData"
// @Success 200 {string} json"{"code": 200,"message": "新增成功"}"
// @Router /api/v1/user [post]
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
// @Router /api/v1/user/{userId} [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "删除用户成功！",
		"data":    user,
	})
}

// GetUserList
// @Summary 获取用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "message"}
// @Router /api/v1/user/list [get]
func GetUserList(c *gin.Context) {
	list := make([]*models.UserBasic, 10)
	list = models.GetUserList()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "获取用户列表",
		"data":    list,
	})
}
