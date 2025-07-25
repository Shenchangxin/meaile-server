package model

import (
	"meaile-server/meaile-user/model"
	"time"
)

type MeaileBookVo struct {
	Id           int64             `gorm:"column:id;primaryKey" json:"id"`           //type:BIGINT       comment:主键        version:2024-11-03 22:16
	BookName     string            `gorm:"column:book_name" json:"bookName"`         //type:string       comment:菜谱名称    version:2024-11-03 22:16
	Image        string            `gorm:"column:image" json:"image"`                //type:string       comment:封面照片    version:2024-11-03 22:16
	ContentMedia string            `gorm:"column:content_media" json:"contentMedia"` //type:string       comment:菜品内容图片/视频OssId   version:2024-08-30 10:24
	Introduction string            `gorm:"column:introduction" json:"introduction"`  //type:string       comment:菜谱介绍    version:2024-11-03 22:16
	Favorite     float64           `gorm:"column:favorite" json:"favorite"`          //type:*float64     comment:评分        version:2024-11-03 22:16
	Sort         int               `gorm:"column:sort" json:"sort"`                  //type:*int         comment:排序        version:2024-11-03 22:16
	Status       string            `gorm:"column:status" json:"status"`              //type:string       comment:状态        version:2024-11-03 22:16
	CreatedBy    string            `gorm:"column:CREATED_BY" json:"createdBy"`       //type:string       comment:创建人      version:2024-11-03 22:16
	CreatedTime  time.Time         `gorm:"column:CREATED_TIME" json:"createdTime"`   //type:*time.Time   comment:创建时间    version:2024-11-03 22:16
	UpdatedBy    string            `gorm:"column:UPDATED_BY" json:"updatedBy"`       //type:string       comment:更新人      version:2024-11-03 22:16
	UpdatedTime  time.Time         `gorm:"column:UPDATED_TIME" json:"updatedTime"`   //type:*time.Time   comment:更新时间    version:2024-11-03 22:16
	TagList      []model.MeaileTag `gorm:"many2many:meaile_book_tag;joinForeignKey:book_id;joinReferences:tag_id;" json:"tagList"`
	ImageOssObj  model.MeaileOss   `gorm:"-" json:"imageOssObj"`
}

func (MeaileBookVo) TableName() string {
	return "meaile_book"
}
