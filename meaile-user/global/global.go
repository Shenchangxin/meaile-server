package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"meaile-server/meaile-user/config"
	"meaile-server/meaile-user/utils"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	MinioClient  *utils.MinioClient
	RedisClient  *redis.Client
)
