package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"ws/config"
)

var DB = Init()

func Init() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	username := config.C.GetString(`mysql.username`)
	password := config.C.GetString(`mysql.password`)
	host := config.C.GetString(`mysql.host`)
	database := config.C.GetString(`mysql.database`)

	connect := username + ":" + password + "@(" + host + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		Logger:         newLogger.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{TablePrefix: "kk_"},
	})
	if err != nil {
		panic(fmt.Sprintf("mysql connect err:%+v", err))
	}

	return DB
}
