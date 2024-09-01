package main

import (
	"flag"
	"fmt"
	"meaile-web/meaile-user/initialize"
	"net"
)

func main() {

	Ip := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 8090, "端口号")

	//初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()

	_, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *Ip, *Port))
	if err != nil {
		panic("服务启动失败")
	}
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic("服务启动失败")
	}

}
