package models

import "gorm.io/gorm"

// Truck model that has unique id as primary key, chassis number and license plate
type Truck struct {
	gorm.Model
	ChassisNumber string `json:"chassisNumber"`
	LicensePlate  string `json:"licensePlate"`
}

type TruckDTO struct {
	ID            uint   `json:"id"`
	ChassisNumber string `json:"chassisNumber"`
	LicensePlate  string `json:"licensePlate"`
}
