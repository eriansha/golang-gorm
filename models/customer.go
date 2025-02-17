package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Address string
}
