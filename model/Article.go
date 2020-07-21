package model

import (
	"github.com/jinzhu/gorm"
)

//Article 文章
type Article struct {
	gorm.Model
	Category Category
	Title    string `gorm:"type:varchar(200);not null" json:"title"`
	Cid      int    `gorm:"type:int;not null" json:"cid"`
	Desc     string `gorm:"type:varchar(400)" json:"desc"`
	Context  string `gorm:"type:longtext" json:"context"`
	Img      string `gorm:"type:varchar(200)" json:"img"`
}
