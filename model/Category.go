package model

import (
	"github.com/jinzhu/gorm"
)

//Category 分类
type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	List int8   `gorm:"type:int;not null" json:"list"`
}
