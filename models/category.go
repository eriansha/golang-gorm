package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string  `gorm:"notNull"`
	IsActive string  `gorm:"default:true"`
	Books    []*Book `gorm:"many2many:book_categories;"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (req *CreateCategoryRequest) ToModel() Category {
	return Category{
		Name: req.Name,
	}
}
