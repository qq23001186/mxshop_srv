package global

import (
	"gorm.io/gorm"
	"nd/order_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
