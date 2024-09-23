package model

import "time"

type MeaileUserFriend struct {
	Id           int64     `gorm:"column:id" json:"id"`                       //type:BIGINT       comment:主键          version:2024-08-23 22:28
	UserIdMain   int64     `gorm:"column:user_id_main" json:"userIdMain"`     //type:BIGINT       comment:主用户id      version:2024-08-23 22:28
	UserIdFriend int64     `gorm:"column:user_id_friend" json:"userIdFriend"` //type:BIGINT       comment:好友用户id    version:2024-08-23 22:28
	GroupId      int64     `gorm:"column:group_id" json:"groupId"`            //type:BIGINT       comment:分组id        version:2024-08-23 22:28
	CreatedBy    string    `gorm:"column:CREATED_BY" json:"createdBy"`        //type:string       comment:创建人        version:2024-08-23 22:28
	CreatedTime  time.Time `gorm:"column:CREATED_TIME" json:"createdTime"`    //type:*time.Time   comment:创建时间      version:2024-08-23 22:28
	UpdatedBy    string    `gorm:"column:UPDATED_BY" json:"updatedBy"`        //type:string       comment:更新人        version:2024-08-23 22:28
	UpdatedTime  time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"`    //type:*time.Time   comment:更新时间      version:2024-08-23 22:28
}
