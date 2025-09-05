package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Data struct {
	DBClient *gorm.DB
}

func NewData() *Data {
	db := MysqlClient()
	return &Data{
		DBClient: db,
	}
}
func MysqlClient() *gorm.DB {
	//连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/NewProject?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //显示sql语句
	})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}

	return db
}
