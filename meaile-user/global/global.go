package global

import (
	"gorm.io/gorm"
	"meaile-server/meaile-user/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
