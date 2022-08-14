package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"nd/goods_srv/config"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
	EsClient     *elastic.Client
)
