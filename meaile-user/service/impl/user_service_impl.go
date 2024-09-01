package impl

import (
	bo "meaile-web/meaile-user/model/bo"
	vo "meaile-web/meaile-user/model/vo"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) SaveUser(userBo bo.MeaileUserBo) bool {

	return true
}

func (u UserServiceImpl) GetUserInfo(userBo bo.MeaileUserBo) vo.MeaileUserVo {
	return vo.MeaileUserVo{}
}

func (u UserServiceImpl) GetUserList(userBo bo.MeaileUserBo) vo.MeaileUserVoList {
	return vo.MeaileUserVoList{}
}

func (u UserServiceImpl) UpdateUser(userBo bo.MeaileUserBo) bool {
	return true
}
