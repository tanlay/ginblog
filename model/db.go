package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db *gorm.DB
	err error
)


func InitDb() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPass,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(utils.DbType,connStr)
	if err != nil {
		fmt.Println("连接数据库失败,请检查参数：",err.Error())
	}

	db.SingularTable(true)		//关闭表名复数形式

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxIdleTime(10*time.Second)

	defer db.Close()
}
