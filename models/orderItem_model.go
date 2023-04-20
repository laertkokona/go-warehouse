package models

import "gorm.io/gorm"

// OrderItem model that has unique id as primary key, item id, order id and quantity
type OrderItem struct {
	gorm.Model
	ItemId   int `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:itemId"`
	OrderId  int `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignkey:roleId"`
	Quantity int `json:"quantity"`
}
