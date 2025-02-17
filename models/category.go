package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"notNull"`
	Description string
	Books       []*Book `gorm:"many2many:book_categories;"`
}
