package models

import "gorm.io/gorm"

// Item model that has unique id as primary key, name, description, unique code, price and category
type Item struct {
	gorm.Model
	Name              string  `json:"name,omitempty"`
	Description       string  `json:"description,omitempty"`
	Code              string  `json:"code,omitempty" gorm:"uniqueIndex;not null"`
	TotalQuantity     int     `json:"totalQuantity,omitempty"`
	AvailableQuantity int     `json:"availableQuantity,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Category          string  `json:"category,omitempty"`
}

type ItemDTO struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Code        string  `json:"code,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
}
