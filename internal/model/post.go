package model

import "gorm.io/gorm"

// Post store post info
type Post struct {
	gorm.Model
	Title   string  `gorm:"column:title;type:varchar(255);not null;default:'';comment:标题"`
	Content *string `gorm:"column:content;type:text;comment:内容"`
}
