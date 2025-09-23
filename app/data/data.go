package data

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//mysql数据库连接配置

type Data struct {
	DBClient *gorm.DB
	RedisDB  *RedisData
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

// redis连接配置

type RedisData struct {
	RedisClient *redis.Client
}

func NewRedisData() *RedisData {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisData{
		RedisClient: rdb,
	}
}

func (r *RedisData) Close() {
	err := r.RedisClient.Close()
	if err != nil {
		panic("redis关闭失败: " + err.Error())
	}
}
