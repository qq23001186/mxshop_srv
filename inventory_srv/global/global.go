package global

import (
	"gorm.io/gorm"
	"nd/inventory_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
