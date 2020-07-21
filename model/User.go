package model

import (
	"github.com/jinzhu/gorm"
)

//User 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` //0是管理员，1是游客
	Iphone   string `gorm:"type:varchar(20)" json:"iphone"`
	Avater   string `gorm:"type:varchar(200)" json:"avater"`
}
