package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string     `gorm:"column:title;type:varchar(255);not null;default:'';comment:标题"`
	Content   *string    `gorm:"column:content;type:text;comment:内容"`
	Slug      string     `gorm:"column:slug;type:varchar(255);unique;comment:文章URL"`
	Status    string     `gorm:"column:status;type:varchar(255);not null;default:'draft';comment:文章状态"`
	Author    string     `gorm:"column:author;type:varchar(255);not null;comment:作者"`
	Category  string     `gorm:"column:category;type:varchar(255);comment:分类"`
}
