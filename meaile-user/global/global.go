package global

import (
	"gorm.io/gorm"
	"meaile-server/meaile-user/config"
	"meaile-server/meaile-user/utils"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	MinioClient  *utils.MinioClient
)
