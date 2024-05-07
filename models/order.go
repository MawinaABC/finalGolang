package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderItem  []Cart  `json:"order_item"`
	TotalPrice float64 `json:"total_price"`
}
