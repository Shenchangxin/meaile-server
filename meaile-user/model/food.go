package model

import "time"

// MeaileFood  菜品信息表。
// 说明:
// 表名:meaile_food
// group: MeaileFood
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.MeaileFood
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.MeaileFood
// version:2024-08-30 10:24
type MeaileFood struct {
	Id           int64     `gorm:"column:id" json:"Id"`                     //type:BIGINT       comment:主键            version:2024-08-30 10:24
	FoodName     string    `gorm:"column:food_name" json:"foodName"`        //type:string       comment:餐品名称        version:2024-08-30 10:24
	Image        string    `gorm:"column:image" json:"image"`               //type:string       comment:菜品封面图片    version:2024-08-30 10:24
	Introduction string    `gorm:"column:introduction" json:"introduction"` //type:string       comment:菜品介绍        version:2024-08-30 10:24
	Favorite     float64   `gorm:"column:favorite" json:"favorite"`         //type:*float64     comment:评分            version:2024-08-30 10:24
	CreatedBy    string    `gorm:"column:CREATED_BY" json:"createdBy"`      //type:string       comment:创建人          version:2024-08-30 10:24
	CreatedTime  time.Time `gorm:"column:CREATED_TIME" json:"createdTime"`  //type:*time.Time   comment:创建时间        version:2024-08-30 10:24
	UpdatedBy    string    `gorm:"column:UPDATED_BY" json:"updatedBy"`      //type:string       comment:更新人          version:2024-08-30 10:24
	UpdatedTime  time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"`  //type:*time.Time   comment:更新时间        version:2024-08-30 10:24
}
