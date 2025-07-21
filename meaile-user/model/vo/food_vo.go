package model

import (
	"meaile-server/meaile-user/model"
	"time"
)

type MeaileFoodVo struct {
	Id            int64             `gorm:"column:id;primaryKey" json:"id"`           //type:BIGINT       comment:主键            version:2024-08-30 10:24
	FoodName      string            `gorm:"column:food_name" json:"foodName"`         //type:string       comment:餐品名称        version:2024-08-30 10:24
	Image         string            `gorm:"column:image" json:"image"`                //type:string       comment:菜品封面图片    version:2024-08-30 10:24
	ContentMedia  string            `gorm:"column:content_media" json:"contentMedia"` //type:string       comment:菜品内容图片/视频OssId   version:2024-08-30 10:24
	Introduction  string            `gorm:"column:introduction" json:"introduction"`  //type:string       comment:菜品介绍        version:2024-08-30 10:24
	Favorite      float64           `gorm:"column:favorite" json:"favorite"`          //type:*float64     comment:评分            version:2024-08-30 10:24
	CreatedBy     string            `gorm:"column:CREATED_BY" json:"createdBy"`       //type:string       comment:创建人          version:2024-08-30 10:24
	CreatedTime   time.Time         `gorm:"column:CREATED_TIME" json:"createdTime"`   //type:*time.Time   comment:创建时间        version:2024-08-30 10:24
	UpdatedBy     string            `gorm:"column:UPDATED_BY" json:"updatedBy"`       //type:string       comment:更新人          version:2024-08-30 10:24
	UpdatedTime   time.Time         `gorm:"column:UPDATED_TIME" json:"updatedTime"`   //type:*time.Time   comment:更新时间        version:2024-08-30 10:24
	Creator       MeaileUserVo      `gorm:"-" json:"creator"`
	TagList       []model.MeaileTag `gorm:"many2many:meaile_food_tag;joinForeignKey:food_id;joinReferences:tag_id;" json:"tagList"`
	ImageOssObj   model.MeaileOss   `gorm:"-" json:"imageOssObj"`
	MediaUrls     []string          `gorm:"-" json:"mediaUrls"`
	UserAvatarUrl string            `gorm:"-" json:"userAvatarUrl"`
	UserName      string            `gorm:"-" json:"userName"`
	IsFollowing   bool              `gorm:"-" json:"isFollowing"`
}

func (MeaileFoodVo) TableName() string {
	return "meaile_food"
}
