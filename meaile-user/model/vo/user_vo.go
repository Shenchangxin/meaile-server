package model

import (
	"meaile-server/meaile-user/model"
	"time"
)

type MeaileUserVo struct {
	Id              int64           `gorm:"column:id" json:"id"`                            //type:BIGINT       comment:                  version:2024-08-01 15:52
	UserName        string          `gorm:"column:user_name" json:"username"`               //type:string       comment:用户名            version:2024-08-01 15:52
	NickName        string          `gorm:"column:nick_name" json:"nickname"`               //type:string       comment:昵称              version:2024-08-01 15:52
	Password        string          `gorm:"column:password" json:"password"`                //type:string       comment:密码              version:2024-08-01 15:52
	Status          string          `gorm:"column:status" json:"status"`                    //type:string       comment:状态              version:2024-08-01 15:52
	Avatar          string          `gorm:"column:avatar" json:"avatar"`                    //type:string       comment:头像              version:2024-08-01 15:52
	BackgroundImage string          `gorm:"column:background_image" json:"backgroundImage"` //type:string       comment:背景图            version:2024-08-01 15:52
	Profile         string          `gorm:"column:profile" json:"profile"`                  //type:string       comment:个人简介          version:2024-08-01 15:52
	Sex             string          `gorm:"column:sex" json:"sex"`                          //type:string       comment:性别（0女1男）    version:2024-08-01 15:52
	Hobby           string          `gorm:"column:hobby" json:"hobby"`                      //type:string       comment:爱好              version:2024-08-01 15:52
	CreatedBy       string          `gorm:"column:CREATED_BY" json:"createdBy"`             //type:string       comment:创建人            version:2024-08-01 15:52
	CreatedTime     time.Time       `gorm:"column:CREATED_TIME" json:"createdTime"`         //type:*time.Time   comment:创建时间          version:2024-08-01 15:52
	UpdatedBy       string          `gorm:"column:UPDATED_BY" json:"updatedBy"`             //type:string       comment:更新人            version:2024-08-01 15:52
	UpdatedTime     time.Time       `gorm:"column:UPDATED_TIME" json:"updatedTime"`         //type:*time.Time   comment:更新时间          version:2024-08-01 15:52
	AvatarOssObj    model.MeaileOss `gorm:"-" json:"avatarOssObj"`
}

type MeaileUserVoList struct {
	total int32
	data  []*MeaileUserVo
}
