package model

import "time"

type MeaileUserBo struct {
	Id              int64      `json:"id"`              //type:BIGINT       comment:                  version:2024-08-01 15:52
	UserName        string     `json:"userName"`        //type:string       comment:用户名            version:2024-08-01 15:52
	NickName        string     `json:"nickName"`        //type:string       comment:昵称              version:2024-08-01 15:52
	Password        string     `json:"password"`        //type:string       comment:密码              version:2024-08-01 15:52
	Status          string     `json:"status"`          //type:string       comment:状态              version:2024-08-01 15:52
	Avatar          string     `json:"avatar"`          //type:string       comment:头像              version:2024-08-01 15:52
	BackgroundImage string     `json:"backgroundImage"` //type:string       comment:背景图            version:2024-08-01 15:52
	Profile         string     `json:"profile"`         //type:string       comment:个人简介          version:2024-08-01 15:52
	Sex             string     `json:"sex"`             //type:string       comment:性别（0女1男）    version:2024-08-01 15:52
	Hobby           string     `json:"hobby"`           //type:string       comment:爱好              version:2024-08-01 15:52
	CreatedBy       string     `json:"createdBy"`       //type:string       comment:创建人            version:2024-08-01 15:52
	CreatedTime     *time.Time `json:"createdTime"`     //type:*time.Time   comment:创建时间          version:2024-08-01 15:52
	UpdatedBy       string     `json:"updatedBy"`       //type:string       comment:更新人            version:2024-08-01 15:52
	UpdatedTime     *time.Time `json:"updatedTime"`     //type:*time.Time   comment:更新时间          version:2024-08-01 15:52
}
