package model

// MeaileUserFollow  用户关注信息表。
import "time"

type MeaileUserFollowBo struct {
	Id             int64     `gorm:"column:primaryKey;id" json:"id"`                //type:BIGINT       comment:主键id          version:2025-04-10 14:28
	UserName       string    `gorm:"column:user_name" json:"userName"`              //type:string       comment:用户名          version:2025-04-10 14:28
	FollowUserName string    `gorm:"column:follow_user_name" json:"followUserName"` //type:string       comment:关注的用户名    version:2025-04-10 14:28
	FollowTime     time.Time `gorm:"column:follow_time" json:"followTime"`          //type:*time.Time   comment:关注时间        version:2025-04-10 14:28
	CreatedBy      string    `gorm:"column:CREATED_BY" json:"createdBy"`            //type:string       comment:创建人          version:2025-04-10 14:28
	CreatedTime    time.Time `gorm:"column:CREATED_TIME" json:"createdTime"`        //type:*time.Time   comment:创建时间        version:2025-04-10 14:28
	UpdatedBy      string    `gorm:"column:UPDATED_BY" json:"updatedBy"`            //type:string       comment:更新人          version:2025-04-10 14:28
	UpdatedTime    time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"`        //type:*time.Time   comment:更新时间        version:2025-04-10 14:28
}

func (MeaileUserFollowBo) TableName() string {
	return "meaile_user_follow"
}
