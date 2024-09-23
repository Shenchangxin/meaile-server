package model

import "time"

type MeaileFriendGroupVo struct {
	Id          int64      `gorm:"column:id" json:"id"`                    //type:BIGINT       comment:主键          version:2024-08-19 17:44
	GroupName   string     `gorm:"column:group_name" json:"groupName"`     //type:string        comment:分组名称      version:2024-08-19 17:44
	UserId      int64      `gorm:"column:user_id" json:"userId"`           //type:BIGINT       comment:所属用户id    version:2024-08-19 17:44
	CreatedBy   string     `gorm:"column:CREATED_BY" json:"createdBy"`     //type:string       comment:创建人        version:2024-08-19 17:44
	CreatedTime *time.Time `gorm:"column:CREATED_TIME" json:"createdTime"` //type:*time.Time   comment:创建时间      version:2024-08-19 17:44
	UpdatedBy   string     `gorm:"column:UPDATED_BY" json:"updatedBy"`     //type:string       comment:更新人        version:2024-08-19 17:44
	UpdatedTime *time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"` //type:*time.Time   comment:更新时间      version:2024-08-19 17:44
}
type FriendGroupListVo struct {
	groupList []MeaileFriendGroupVo
}
