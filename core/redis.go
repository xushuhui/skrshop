package core

import "github.com/go-redis/redis"

// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接
func InitRedis() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = RDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
