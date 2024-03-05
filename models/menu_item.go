package models

type MenuItem struct {
	ID           int64   `gorm:"primaryKey"`
	Name         string  `gorm:"not null"`
	Price        float64 `gorm:"not null"`
	Description  string
	RestaurantID int64
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
}
