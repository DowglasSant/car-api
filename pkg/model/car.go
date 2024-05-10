package model

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	ID           int     `gorm:"column:id"`
	Year         int     `gorm:"column:year"`
	Make         string  `gorm:"column:make"`
	CarModel     string  `gorm:"column:model"`
	Color        string  `gorm:"column:color"`
	LicensePlate string  `gorm:"column:license_plate"`
	Mileage      int     `gorm:"column:mileage"`
	Price        float64 `gorm:"column:price"`
}
