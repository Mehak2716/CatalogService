package models

type Location struct {
	ID           int64   `gorm:"primaryKey"`
	XCordinate   float64 `gorm:"not null"`
	YCordinate   float64 `gorm:"not null"`
	RestaurantID uint
}
