package model

import "time"

type MeaileFoodTag struct {
	Id          int64     `gorm:"column:primaryKey;id" json:"Id"`         //type:BIGINT       comment:主键        version:2024-08-30 10:55
	FoodId      int64     `gorm:"column:food_id" json:"FoodId"`           //type:BIGINT       comment:菜品id      version:2024-08-30 10:55
	Tag         string    `gorm:"column:tag" json:"Tag"`                  //type:varchar      comment:标签        version:2024-08-30 10:55
	Status      string    `gorm:"column:status" json:"Status"`            //type:string       comment:状态        version:2024-08-30 10:55
	CreatedBy   string    `gorm:"column:CREATED_BY" json:"CreatedBy"`     //type:string       comment:创建人      version:2024-08-30 10:55
	CreatedTime time.Time `gorm:"column:CREATED_TIME" json:"CreatedTime"` //type:*time.Time   comment:创建时间    version:2024-08-30 10:55
}
