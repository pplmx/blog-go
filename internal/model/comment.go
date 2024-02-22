package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string `gorm:"column:content;type:varchar(255);not null;default:'';comment:评论内容"`
	PostID   uint   `gorm:"column:post_id;not null;index;comment:评论的文章ID"`
	ParentID uint   `gorm:"column:parent_id;comment:父评论ID"`
	Rating   int    `gorm:"column:rating;type:int;comment:评论评分"`
}
