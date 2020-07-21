package model

import (
	"fmt"
	"ginblog/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //驱动
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//InitDB 初始化数据库
func InitDB() {
	db, err := gorm.Open(utils.DB, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DBUsername,
		utils.DBPassWord,
		utils.DBHost,
		utils.DBPort,
		utils.DBName,
	))
	if err != nil {
		log.Println("连接数据库失败，请检查参数", err)
	}
	// defer db.Close()
	//禁用默认表名的复数形式
	db.SingularTable(true)

	//自动数据表迁移
	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(10 * time.Second)

}
