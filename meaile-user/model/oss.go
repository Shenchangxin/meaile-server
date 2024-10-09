package model

import "time"

type MeaileOss struct {
	Id          int64     `gorm:"column:primaryKey;id" json:"id"`         //type:BIGINT       comment:主键id        version:2024-9-09 22:30
	OssId       string    `gorm:"column:oss_id" json:"ossId"`             //type:string       comment:桶存储名称    version:2024-9-09 22:30
	FileName    string    `gorm:"column:file_name" json:"fileName"`       //type:string       comment:文件名称      version:2024-9-09 22:30
	Suffix      string    `gorm:"column:suffix" json:"suffix"`            //type:string       comment:文件后缀      version:2024-9-09 22:30
	FileUrl     string    `gorm:"column:file_url" json:"fileUrl"`         //type:string       comment:文件地址      version:2024-9-09 22:30
	CreatedBy   string    `gorm:"column:CREATED_BY" json:"createdBy"`     //type:string       comment:创建人        version:2024-9-09 22:30
	CreatedTime time.Time `gorm:"column:CREATED_TIME" json:"createdTime"` //type:*time.Time   comment:创建时间      version:2024-9-09 22:30
	UpdatedBy   string    `gorm:"column:UPDATED_BY" json:"updatedBy"`     //type:string       comment:更新人        version:2024-9-09 22:30
	UpdatedTime time.Time `gorm:"column:UPDATED_TIME" json:"updatedTime"` //type:*time.Time   comment:更新时间      version:2024-9-09 22:30
}
