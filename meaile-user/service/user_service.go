package service

import (
	bo "meaile-web/meaile-user/model/bo"
	vo "meaile-web/meaile-user/model/vo"
)

type UserService interface {
	GetUserInfo(userBo bo.MeaileUserBo) vo.MeaileUserVo
	GetUserList(userBo bo.MeaileUserBo) vo.MeaileUserVoList
	SaveUser(userBo bo.MeaileUserBo) bool
	UpdateUser(userBo bo.MeaileUserBo) bool
}
