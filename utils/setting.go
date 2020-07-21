package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

//全局变量
var (
	AppMode  string
	HTTPPort string

	DB         string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassWord string
	DBName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadDatabase(file)
}

//LoadServer 服务配置文件解析
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HTTPPort = file.Section("server").Key("HttpPort").MustString(":9999")
}

//LoadDatabase DB配置文件解析
func LoadDatabase(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBUsername = file.Section("database").Key("DBUsername").MustString("ginblog")
	DBPassWord = file.Section("database").Key("DBPassWord").MustString("admin")
	DBName = file.Section("database").Key("DBName").MustString("ginblog")
}
