package impl

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	"meaile-server/meaile-user/model"
	bo "meaile-server/meaile-user/model/bo"
	vo "meaile-server/meaile-user/model/vo"
	"time"
)

type UserServiceImpl struct {
}

func (u *UserServiceImpl) SaveUser(ctx *gin.Context, userBo bo.MeaileUserBo) (bool, error) {
	var user model.MeaileUser
	result := global.DB.Where(&model.MeaileUser{
		UserName: userBo.UserName,
	}).First(&user)
	if result.RowsAffected == 1 {
		//model.Success(ctx, model.FAILED, nil)
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

func (u *UserServiceImpl) GetUserInfo(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVo {
	return vo.MeaileUserVo{}
}

func (u *UserServiceImpl) GetUserList(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVoList {
	return vo.MeaileUserVoList{}
}

func (u *UserServiceImpl) UpdateUser(ctx *gin.Context, userBo bo.MeaileUserBo) bool {
	return true
}

func (u *UserServiceImpl) Login(ctx *gin.Context, loginBo bo.LoginForm) *model.Response {
	var user model.MeaileUser
	result := global.DB.Where(&model.MeaileUser{
		UserName: loginBo.UserName,
	}).First(&user)

	if result.RowsAffected == 1 {
		//校验密码是否正确
		pass, err := CheckPassword(user.Password, loginBo.Password)
		if err != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "登录失败",
				Data: err,
			}
		} else {
			if pass {

				j := middlewares.NewJWT()
				claims := model.CustomClaims{
					ID:          uint(user.Id),
					NickName:    user.NickName,
					AuthorityId: 1,
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),
						ExpiresAt: time.Now().Unix() + 60*60*24*30,
						Issuer:    "meaile",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					return &model.Response{
						Code: model.FAILED,
						Msg:  "登录失败",
						Data: err,
					}
				}
				return &model.Response{
					Code: model.SUCCESS,
					Msg:  "登录成功",
					Data: map[string]interface{}{
						"id":         user.Id,
						"nick_name":  user.NickName,
						"token":      token,
						"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
					},
				}
				//ctx.JSON(http.StatusOK, gin.H{
				//	"id":         user.Id,
				//	"nick_name":  user.NickName,
				//	"token":      token,
				//	"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				//})
			} else {
				return &model.Response{
					Code: model.FAILED,
					Msg:  "登录失败,密码不正确",
					Data: nil,
				}
			}
		}
	} else {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "登录失败,未找到该用户",
			Data: nil,
		}
	}
}

func (u *UserServiceImpl) Register(ctx *gin.Context, registerUserBo bo.MeaileUserBo) *model.Response {
	var user model.MeaileUser
	result := global.DB.Where(&model.MeaileUser{
		UserName: registerUserBo.UserName,
	}).First(&user)

	if result.RowsAffected == 1 {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "用户名重复",
			Data: nil,
		}
	} else {
		user.Sex = registerUserBo.Sex
		user.UserName = registerUserBo.UserName
		user.NickName = registerUserBo.NickName
		user.Status = "0"
		user.Avatar = registerUserBo.Avatar
		user.BackgroundImage = registerUserBo.BackgroundImage
		user.Hobby = registerUserBo.Hobby
		user.Profile = registerUserBo.Profile
		user.CreatedBy = user.UserName
		user.CreatedTime = time.Now()
		//salt, _ := bcrypt.GenerateFromPassword([]byte("shenchangxin"), bcrypt.DefaultCost)
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserBo.Password), bcrypt.DefaultCost)
		if err != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "创建失败",
				Data: err,
			}
		}
		user.Password = string(encryptedPassword)
		if result.Error != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "创建失败",
				Data: result.Error,
			}
		}
		result = global.DB.Create(&user)
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "注册成功",
			Data: user,
		}
	}
}

func CheckPassword(passwordDB string, passwordLogin string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordLogin))
	if err != nil {
		return false, err
	} else {
		return false, nil
	}
}
