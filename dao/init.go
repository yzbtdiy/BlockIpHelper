package dao

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/yzbtdiy/BlockIpHelper/models"
)

var db *gorm.DB

// 初始化数据库
func InitDb() {
	var err error
	db, err = gorm.Open(sqlite.Open("./data/app.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(1)
	tableInit(&models.WhiteListTable{})
	// tableInit(&models.BlackListTable{})
}

// 检查表是否存在, 不存在则创建
func tableInit(table interface{}) {
	if db.Migrator().HasTable(table) {
		return
	} else {
		db.Migrator().CreateTable(table)
	}
}
