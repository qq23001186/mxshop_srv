package global

import (
	"gorm.io/gorm"
	"nd/user_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
)
