package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ID 			uint 	`gorm:"primary_key;unique" json:"id"`
	Name 		string 	`json:"name" gorm:"not null" validate:"required"`
	Price 		float64 `json:"price" gorm:"not null" validate:"required"`
	CategoryID 	uint 	`json:"category_id"`
}