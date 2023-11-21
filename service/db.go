package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"xapi/model"
)

var (
	db  *gorm.DB
	rdb *redis.Client
)

func GetDB() *gorm.DB {
	if db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			GetConfig().Mysql.UserName,
			GetConfig().Mysql.Password,
			GetConfig().Mysql.Host,
			GetConfig().Mysql.Port,
			GetConfig().Mysql.Database)

		db_tmp, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: gormlogger.Default.LogMode(gormlogger.Info),
		})
		if err != nil {
			panic("failed to connect database")
		}
		log.Println("数据库初始化成功...")
		db = db_tmp
		createTable(db_tmp)
	}
	return db
}

func createTable(database *gorm.DB) {
	if err := database.AutoMigrate(
		&model.User{},
	); err != nil {
		log.Println("建表时出现错误", err)
	}

	log.Println("建表成功...")
}
