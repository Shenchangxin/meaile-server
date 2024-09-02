package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"meaile-server/meaile-user/global"
)

func InitConfig() {

	//debug := GetEnvInfo("GO-BLOG")
	configFilePrefix := "config"

	configFileName := fmt.Sprintf("meaile-user/%s-dev.yaml", configFilePrefix)
	//if debug {
	//	configFileName = fmt.Sprintf("user-srv/%s-dev.yaml", configFilePrefix)
	//}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		panic(err)
	}
	zap.S().Infof("配置信息：%v", global.ServerConfig)

}
