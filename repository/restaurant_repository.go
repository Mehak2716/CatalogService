package repository

import (
	"catalog/models"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	DB *gorm.DB
}

func (repo *RestaurantRepository) Save(restaurant *models.Restaurant) (*models.Restaurant, error) {

	res := repo.DB.Create(restaurant)
	if res.Error != nil {
		return nil, res.Error
	}
	return restaurant, nil
}

func (repo *RestaurantRepository) FetchAll() ([]models.Restaurant, error) {

	var restaurants []models.Restaurant
	res := repo.DB.Preload("Location").Find(&restaurants)

	if res.Error != nil {
		return nil, res.Error
	}

	return restaurants, nil
}

func (repo *RestaurantRepository) IsExists(restaurant models.Restaurant) bool {
	var count int64
	repo.DB.Model(&models.Restaurant{}).
		Joins("JOIN locations ON restaurants.id = locations.restaurant_id").
		Where("restaurants.name = ? AND locations.x_cordinate = ? AND locations.y_cordinate = ? AND restaurants.deleted_at IS NULL", restaurant.Name, restaurant.Location.XCordinate, restaurant.Location.YCordinate).
		Count(&count)

	return count > 0
}
