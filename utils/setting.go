package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string

	DbType string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径: ", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadDatabase(file *ini.File) {
	DbType = file.Section("database").Key("DbType").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPass = file.Section("database").Key("DbPass").MustString("admin@123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}