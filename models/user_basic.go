package models

import (
	"time"

	"gorm.io/gorm"

	"go-chat/global/orm"
)

type UserBasic struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"name"`
	PassWord      string `json:"password"`
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	LoginOutTime  time.Time
	HeartbeatTime time.Time
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	orm.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return orm.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return orm.DB.Delete(&user)
}

func GetUserList() []*UserBasic {
	list := make([]*UserBasic, 10)
	orm.DB.Find(&list)

	return list
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	orm.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}
