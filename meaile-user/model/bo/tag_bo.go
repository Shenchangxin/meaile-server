package model

import "time"

// MeaileTag  标签信息表。
type MeaileTagBo struct {
	Id          int64      `gorm:"column:id" json:"Id"`                    //type:BIGINT       comment:主键          version:2024-10-03 00:18
	TagName     string     `gorm:"column:tag_name" json:"tagName"`         //type:string       comment:标签名称      version:2024-10-03 00:18
	ParentId    int64      `gorm:"column:parent_id" json:"parentId"`       //type:BIGINT       comment:父标签id      version:2024-10-03 00:18
	UserId      int64      `gorm:"column:user_id" json:"userId"`           //type:BIGINT       comment:所属用户id    version:2024-10-03 00:18
	Status      string     `gorm:"column:status" json:"status"`            //type:string       comment:状态          version:2024-10-03 00:18
	CreatedBy   string     `gorm:"column:CREATED_BY" json:"createdBy"`     //type:string       comment:创建人        version:2024-10-03 00:18
	CreatedTime *time.Time `gorm:"column:CREATED_TIME" json:"createdTime"` //type:*time.Time   comment:创建时间      version:2024-10-03 00:18
	UpdatedBy   string     `gorm:"column:UPDATED_BY" json:"updatedBy"`     //type:string       comment:更新人        version:2024-10-03 00:18
	UpdatedTime *time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"` //type:*time.Time   comment:更新时间      version:2024-10-03 00:18
}
