package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB 初始化数据库（只调用一次）
func InitDB() {
	once.Do(func() {
		dsn := "root:csh020318.@tcp(127.0.0.1:3306)/apple_traceability_system?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("连接数据库失败:", err)
		}

		fmt.Println("数据库连接成功")
	})
}

// GetDB 获取全局数据库连接
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
