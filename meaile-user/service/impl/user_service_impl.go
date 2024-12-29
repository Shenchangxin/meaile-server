package impl

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"meaile-server/meaile-user/global"
	"meaile-server/meaile-user/middlewares"
	_ "meaile-server/meaile-user/middlewares"
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
		CreatedBy:       userBo.UserName,
		CreatedTime:     time.Now(),
		UpdatedBy:       userBo.UserName,
		UpdatedTime:     time.Now(),
	}
	result = global.DB.Create(user)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u *UserServiceImpl) GetUserList(ctx *gin.Context, userBo bo.MeaileUserBo) vo.MeaileUserVoList {
	return vo.MeaileUserVoList{}
}

func (u *UserServiceImpl) UpdateUser(ctx *gin.Context, registerUserBo bo.MeaileUserBo) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err,
		}
	}
	var user model.MeaileUser
	result := global.DB.Where(&model.MeaileUser{
		UserName: registerUserBo.UserName,
	}).First(&user)
	if result.RowsAffected == 1 {
		user.Sex = registerUserBo.Sex
		user.UserName = registerUserBo.UserName
		user.NickName = registerUserBo.NickName
		user.Status = "0"
		user.Avatar = registerUserBo.Avatar
		user.BackgroundImage = registerUserBo.BackgroundImage
		user.Hobby = registerUserBo.Hobby
		user.Profile = registerUserBo.Profile
		user.CreatedBy = customClaims.UserName
		user.UpdatedTime = time.Now()
		result = global.DB.Save(&user)
		if result.Error != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "修改用户信息失败",
				Data: result.Error,
			}
		}
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "操作成功",
			Data: user,
		}
	} else {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "未找到该用户信息",
			Data: nil,
		}
	}
}
func (u *UserServiceImpl) GetUserFriendList(ctx *gin.Context, token string) *model.Response {
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err.Error(),
		}
	}
	result := make(map[string][]model.MeaileUser)
	var rawResults []struct {
		GroupName string
		ID        uint
		UserName  string
		Avatar    string
		NickName  string
	}
	var relations []struct {
		GroupName string
		Friend    model.MeaileUser
	}

	err = global.DB.Table("meaile_user_friend muf").
		Select("mfg.group_name as GroupName, mu.id as ID,mu.avatar as Avatar, mu.user_name as UserName,mu.nick_name as NickName").
		Joins("join meaile_friend_group mfg on muf.group_id = mfg.id").
		Joins("join meaile_user mu on muf.user_id_friend = mu.id").
		Where("muf.user_id_main = ?", customClaims.ID).
		Scan(&rawResults).Error
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "查询失败",
			Data: err.Error(),
		}
	}
	for _, rawResult := range rawResults {
		relations = append(relations, struct {
			GroupName string
			Friend    model.MeaileUser
		}{
			GroupName: rawResult.GroupName,
			Friend: model.MeaileUser{
				Id:       int64(rawResult.ID),
				UserName: rawResult.UserName,
				NickName: rawResult.NickName,
				Avatar:   rawResult.Avatar,
			},
		})
	}
	for _, relation := range relations {
		result[relation.GroupName] = append(result[relation.GroupName], relation.Friend)
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: result,
	}
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
				Data: err.Error(),
			}
		} else {
			if pass {

				j := middlewares.NewJWT()
				claims := model.CustomClaims{
					ID:          uint(user.Id),
					UserName:    user.UserName,
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
						Data: err.Error(),
					}
				}
				// 将日志写入数据库
				tx := global.DB.Exec("INSERT INTO meaile_login_log (id,login_time, login_ip, login_user) VALUES ( ?,?, ?,?)", nil, time.Now(), ctx.ClientIP(), user.UserName)
				if tx.Error != nil {
					zap.S().Errorw("无法插入登录日志: %v\n", tx.Error)
					return &model.Response{
						Code: model.FAILED,
						Msg:  "无法插入登录日志",
						Data: tx.Error,
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
		user.UpdatedTime = time.Now()
		user.UpdatedBy = user.UserName
		user.CreatedTime = time.Now()
		//salt, _ := bcrypt.GenerateFromPassword([]byte("shenchangxin"), bcrypt.DefaultCost)
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserBo.Password), bcrypt.DefaultCost)
		if err != nil {
			return &model.Response{
				Code: model.FAILED,
				Msg:  "创建失败",
				Data: err.Error(),
			}
		}
		user.Password = string(encryptedPassword)

		result = global.DB.Create(&user)
		return &model.Response{
			Code: model.SUCCESS,
			Msg:  "注册成功",
			Data: user,
		}
	}
}

func (u *UserServiceImpl) GetUserInfo(ctx *gin.Context, token string) *model.Response {
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err.Error(),
		}
	}
	var user model.MeaileUser
	result := global.DB.Where("user_name = ?", customClaims.UserName).First(&user)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: result.Error.Error(),
		}
	}
	var avatarOss model.MeaileOss
	result = global.DB.Where("oss_id = ?", user.Avatar).First(&avatarOss)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: result.Error.Error(),
		}
	}
	fileUrl, _ := global.MinioClient.GetPresignedGetObject(global.ServerConfig.MinioConfig.BucketName, avatarOss.OssId+avatarOss.Suffix, 24*time.Hour)
	avatarOss.FileUrl = fileUrl
	user.AvatarOssObj = avatarOss
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "获取用户信息成功",
		Data: user,
	}
}

func (u *UserServiceImpl) AddFriend(ctx *gin.Context, addFriendBo bo.AddUserFriendBo) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	if token == "" {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "登录过期，请重新登录",
			Data: nil,
		}
	}
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err.Error(),
		}
	}
	var userFriends []model.MeaileUserFriend
	for _, userId := range addFriendBo.UserIds {
		userFriend := model.MeaileUserFriend{
			UserIdFriend: userId,
			UserIdMain:   int64(customClaims.ID),
			GroupId:      addFriendBo.GroupId,
			CreatedBy:    customClaims.UserName,
			CreatedTime:  time.Now(),
			UpdatedTime:  time.Now(),
			UpdatedBy:    customClaims.UserName,
		}
		userFriends = append(userFriends, userFriend)
	}
	_ = global.DB.Where("user_id_main = ? and user_id_friend in (?)", customClaims.ID, addFriendBo.UserIds).Delete(&model.MeaileUserFriend{})
	result := global.DB.Create(&userFriends)
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "操作失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: userFriends,
	}
}
func (u *UserServiceImpl) DeleteFriend(ctx *gin.Context, userId int64) *model.Response {
	token := ctx.Request.Header.Get("x-token")
	if token == "" {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "登录过期，请重新登录",
			Data: nil,
		}
	}
	myJwt := middlewares.NewJWT()
	customClaims, err := myJwt.ParseToken(token)
	if err != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "获取用户信息失败，请重新登录",
			Data: err.Error(),
		}
	}
	result := global.DB.Where("user_id_main = ? and user_id_friend = ?", customClaims.ID, userId).Delete(&model.MeaileUserFriend{})
	if result.Error != nil {
		return &model.Response{
			Code: model.FAILED,
			Msg:  "删除失败",
			Data: result.Error,
		}
	}
	return &model.Response{
		Code: model.SUCCESS,
		Msg:  "操作成功",
		Data: nil,
	}
}
func CheckPassword(passwordDB string, passwordLogin string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(passwordLogin))
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
