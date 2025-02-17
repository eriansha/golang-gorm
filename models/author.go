package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name        string
	Description string
}

type CreateAuthorRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (req *CreateAuthorRequest) ToModel() Author {
	return Author{
		Name:        req.Name,
		Description: req.Description,
	}
}
