package global

import (
	"gorm.io/gorm"
	"meaile-web/meaile-user/config"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
)
