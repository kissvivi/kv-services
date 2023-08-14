package data

import (
	"fmt"
	"kv-services/config"
	"kv-services/db"
)

func InitDB(cfg *config.Config) {
	//初始化数据库
	baseDB := db.NewBaseDB("mysql")
	fmt.Println(cfg)
	baseDB.InitDB(cfg)
	db.MYSQLDB.AutoMigrate(&User{}, Module{}, Role{}, UserRole{}, Permission{})
}
