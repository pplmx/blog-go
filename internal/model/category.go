package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(255);not null;unique;comment:分类名称"`
	Description string `gorm:"column:description;type:varchar(255);comment:分类描述"`
	Slug        string `gorm:"column:slug;type:varchar(255);unique;comment:分类URL"`
	Count       int    `gorm:"column:count;type:int;comment:分类下的文章数"`
	Posts       []Post `gorm:"many2many:post_categories;comment:分类下的文章"`
}
