package main

import (
	"BackEnd/router"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:csh020318.@tcp(127.0.0.1:3306)/apple_traceability_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := InitDB()
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	r := router.SetupRouter(db)

	fmt.Println("服务已启动：http://localhost:8080")
	r.Run(":8080")
}
