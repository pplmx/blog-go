package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(255);not null;unique;comment:标签名称"`
	Description string `gorm:"column:description;type:varchar(255);comment:标签描述"`
	Slug        string `gorm:"column:slug;type:varchar(255);unique;comment:标签URL"`
	Count       int    `gorm:"column:count;type:int;comment:标签下的文章数"`
	Posts       []Post `gorm:"many2many:post_tags;comment:标签下的文章"`
}
