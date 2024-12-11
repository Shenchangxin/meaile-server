package initialize

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"meaile-server/meaile-user/global"
)

func InitRedis() {
	redisConfig := global.ServerConfig.RedisConfig
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port), // Redis地址
		Username: redisConfig.Username,
		Password: redisConfig.Password, // Redis密码，如果没有则为空字符串
		DB:       redisConfig.Database, // 使用默认DB
	})
	global.RedisClient = redisClient
}
