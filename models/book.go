package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"notNull"`
	Description string
	Year        int16
	AuthorID    int
	Author      Author
	CategoryID  int
	Categories  []*Category `gorm:"many2many:book_categories;"`
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Year        int16  `json:"year"`
	AuthorID    int    `json:"authorId" binding:"required"`
}

func (req *CreateBookRequest) ToModel() Book {
	return Book{
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		AuthorID:    req.AuthorID,
	}
}
