package model

import "time"

type MeaileBookTag struct {
	Id          int64     `gorm:"column:id" json:"id"`                    //type:BIGINT       comment:主键id      version:2024-11-04 21:19
	BookId      int64     `gorm:"column:book_id" json:"bookId"`           //type:BIGINT       comment:菜谱id      version:2024-11-04 21:19
	TagId       int64     `gorm:"column:tag_id" json:"tagId"`             //type:BIGINT       comment:标签id      version:2024-11-04 21:19
	Status      string    `gorm:"column:status" json:"status"`            //type:string       comment:状态        version:2024-11-04 21:19
	CreatedBy   string    `gorm:"column:CREATED_BY" json:"createdBy"`     //type:string       comment:创建人      version:2024-11-04 21:19
	CreatedTime time.Time `gorm:"column:CREATED_TIME" json:"createdTime"` //type:*time.Time   comment:创建时间    version:2024-11-04 21:19
	UpdatedBy   string    `gorm:"column:UPDATED_BY" json:"updatedBy"`     //type:string       comment:更新人      version:2024-11-04 21:19
	UpdatedTime time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"` //type:*time.Time   comment:更新时间    version:2024-11-04 21:19
}
