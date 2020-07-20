package core

import (
	"github.com/go-redis/redis"
	"log"
	"skrshop-api/config"
)

// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接
func initRedis() {
	NewClient()
	_, err := RDB.Ping().Result()
	if err != nil {
		log.Fatalf(" [ERROR] redis connect :%s err:%v\n", RDB, err)
	}

}
func NewClient() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: config.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})

}
