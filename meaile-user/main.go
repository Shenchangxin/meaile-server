package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/initialize"
	"net"
)

func main() {

	Ip := flag.String("ip", "0.0.0.0", "ip地址")
	//Port := flag.Int("port", 8090, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()
	Routers := initialize.Routers()

	zap.S().Infof("启动服务器：%d", global.ServerConfig.Port)

	if err := Routers.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panicf("启动服务器：%d 失败", global.ServerConfig.Port)
	}

	_, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *Ip, global.ServerConfig.Port))
	if err != nil {
		panic("服务启动失败")
	}
}
