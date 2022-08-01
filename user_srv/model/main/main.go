package main

import (
	"fmt"
	"log"
	"nd/user_srv/global"
	"nd/user_srv/initialize"
	"nd/user_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	initialize.InitConfig()
	dsn := fmt.Sprintf("root:jiushi@tcp(%s:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.Host)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	// 迁移 schema
	_ = db.AutoMigrate(&model.User{}) //此处应该有sql语句
}
