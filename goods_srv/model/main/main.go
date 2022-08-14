package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"nd/goods_srv/global"
	"nd/goods_srv/initialize"
	"nd/goods_srv/model"
	"os"
	"strconv"
	"time"
)

func main() {
	/*	initialize.InitConfig()
		dsn := fmt.Sprintf("root:jiushi@tcp(%s:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.Host)

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

		_ = db.AutoMigrate(&model.Category{},
			&model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})*/
	Mysql2Es()
}

// Mysql2Es 将之前mysql的goods数据保存到es中
func Mysql2Es() {
	initialize.InitConfig()
	dsn := fmt.Sprintf("root:jiushi@tcp(%s:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local", global.ServerConfig.MysqlInfo.Host)

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

	host := fmt.Sprintf("http://%s:%d", global.ServerConfig.EsInfo.Host, global.ServerConfig.EsInfo.Port)
	loggerEs := log.New(os.Stdout, "mxshop", log.LstdFlags)
	global.EsClient, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false),
		elastic.SetTraceLog(loggerEs))
	if err != nil {
		panic(err)
	}

	var goods []model.Goods
	db.Find(&goods)
	for _, g := range goods {
		esModel := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}
		_, err = global.EsClient.Index().
			Index(esModel.GetIndexName()).
			BodyJson(esModel).
			Id(strconv.Itoa(int(g.ID))).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
		//强调一下 一定要将docker启动es的java_ops的内存设置大一些 否则运行过程中会出现 bad request错误
	}
}
