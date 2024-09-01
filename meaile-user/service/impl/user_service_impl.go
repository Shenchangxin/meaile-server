package impl

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"meaile-web/meaile-user/global"
	"meaile-web/meaile-user/model"
	bo "meaile-web/meaile-user/model/bo"
	vo "meaile-web/meaile-user/model/vo"
	"time"
)

type UserServiceImpl struct {
}

func (u UserServiceImpl) SaveUser(ctx *gin.Context, userBo bo.MeaileUserBo) (bool, error) {
	var user model.MeaileUser
	result := global.DB.Where(&model.MeaileUser{
		UserName: userBo.UserName,
	}).First(&user)
	if result.RowsAffected == 1 {
		model.Success(ctx, model.SUCCESS, nil)
		return false, nil
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userBo.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user = model.MeaileUser{
		UserName:        userBo.UserName,
		Password:        string(encryptedPassword),
		NickName:        userBo.NickName,
		Status:          "0",
		Avatar:          userBo.Avatar,
		BackgroundImage: userBo.BackgroundImage,
		Profile:         userBo.Profile,
		Sex:             userBo.Sex,
		Hobby:           userBo.Hobby,
		CreatedBy:       "",
		CreatedTime:     time.Now(),
		UpdatedBy:       "",
		UpdatedTime:     time.Now(),
	}
	result = global.DB.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u UserServiceImpl) GetUserInfo(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVo {
	return vo.MeaileUserVo{}
}

func (u UserServiceImpl) GetUserList(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVoList {
	return vo.MeaileUserVoList{}
}

func (u UserServiceImpl) UpdateUser(ctx *gin.Context, userBo bo.MeaileUserBo) bool {
	return true
}
