package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Category    string  `gorm:"unique" json:"category"`
	Description string  `json:"description" type:"text"`
	Price       float64 `json:"price"`
}
