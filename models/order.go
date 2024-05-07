package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderItem  []Cart  `json:"order_item"`
	UserName   string  `json:"user_name"`
	TotalPrice float64 `json:"total_price"`
	Status     int     `json:"status"`
}
