package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name       string
	PassWord   string
	Phone      string
	Email      string
	Identity   string
	ClientIp   string
	ClientPort string
	LoginTime  time.Time
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
