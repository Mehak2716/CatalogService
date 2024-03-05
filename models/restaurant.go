package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name     string   `gorm:"not null"`
	Location Location `gorm:"foreignKey:RestaurantID;constraint:OnDelete:CASCADE"`
	Status   string
}
