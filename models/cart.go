package models

import "gorm.io/gorm"

type Cart struct {
	UserId       uint    `json:"user_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	Status       int     `json:"status"`
	gorm.Model   `json:"gorm.Model"`
}
