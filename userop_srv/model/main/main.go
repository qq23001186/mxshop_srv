package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"nd/userop_srv/global"
	"nd/userop_srv/initialize"
	"nd/userop_srv/model"
	"os"
	"time"
)

func main() {
	initialize.InitConfig()
	dsn := fmt.Sprintf("root:jiushi@tcp(%s:3306)/mxshop_userop_srv?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.Host)

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

	_ = db.AutoMigrate(&model.LeavingMessages{}, &model.UserFav{}, model.Address{})
}
