package initialize

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/utils"
)

func InitHuaWeiOBSClient() {

	ak := global.ServerConfig.HuaWeiOBSConfig.AccessKey
	sk := global.ServerConfig.HuaWeiOBSConfig.SecretKey
	// endpoint填写Bucket对应的Endpoint, 这里以华北-北京四为例，其他地区请按实际情况填写。
	endPoint := global.ServerConfig.HuaWeiOBSConfig.EndPoint
	// 创建obsClient实例
	obsClient, err := obs.New(ak, sk, endPoint)
	if err == nil {
		// 使用访问OBS

		// 关闭obsClient
		obsClient.Close()
	}
	client := &utils.HuaWeiOBSClient{
		OBSClient: obsClient,
	}
	global.HuaWeiOBSClient = client

}
