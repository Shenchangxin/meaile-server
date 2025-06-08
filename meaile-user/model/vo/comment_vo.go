package model

import "time"

// MeaileComment  评论表。
type MeaileCommentVo struct {
	Id             int64        `gorm:"column:id" json:"id"`                          //type:BIGINT       comment:评论主键id                 version:2025-05-05 22:08
	UserId         int64        `gorm:"column:user_id" json:"userId"`                 //type:BIGINT       comment:用户id                     version:2025-05-05 22:08
	CommentContent string       `gorm:"column:comment_content" json:"commentContent"` //type:VARCHAR      comment:评论内容                   version:2025-05-05 22:08
	BizId          int64        `gorm:"column:biz_id" json:"bizId"`                   //type:BIGINT       comment:业务id（菜品id/菜谱id）    version:2025-05-05 22:08
	FatherId       int64        `gorm:"column:father_id" json:"fatherId"`             //type:BIGINT       comment:父id，0表示一级评论        version:2025-05-05 22:08
	CreatedBy      string       `gorm:"column:CREATED_BY" json:"createdBy"`           //type:VARCHAR      comment:创建人                     version:2025-05-05 22:08
	CreatedTime    time.Time    `gorm:"column:CREATED_TIME" json:"createdTime"`       //type:*time.Time   comment:创建时间                   version:2025-05-05 22:08
	UpdatedBy      string       `gorm:"column:UPDATED_BY" json:"updatedBy"`           //type:VARCHAR      comment:更新人                     version:2025-05-05 22:08
	UpdatedTime    time.Time    `gorm:"column:UPDATED_TIME" json:"updatedTime"`       //type:*time.Time   comment:更新时间                   version:2025-05-05 22:08
	Creator        MeaileUserVo `gorm:"-" json:"creator"`
}

// TableName 表名:meaile_comment，评论表。
// 说明:
func (MeaileCommentVo) TableName() string {
	return "meaile_comment"
}
