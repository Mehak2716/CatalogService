package repository

import (
	"catalog/models"

	"gorm.io/gorm"
)

type MenuItemRepository struct {
	DB *gorm.DB
}

func (repo *MenuItemRepository) Save(menuItem *models.MenuItem) (*models.MenuItem, error) {
	res := repo.DB.Create(menuItem)

	if res.Error != nil {
		return nil, res.Error
	}

	return menuItem, nil
}

func (repo *MenuItemRepository) FetchAll(restaurantID int) ([]models.MenuItem, error) {

	var menuItems []models.MenuItem
	res := repo.DB.Where("restaurant_id = ?", restaurantID).Find(&menuItems)

	if res.Error != nil {
		return nil, res.Error
	}

	return menuItems, nil
}

func (repo *MenuItemRepository) IsExists(menuItem models.MenuItem) bool {
	var count int64
	repo.DB.Model(&models.MenuItem{}).
		Where("name = ? AND price = ? AND description = ? AND restaurant_id = ?", menuItem.Name, menuItem.Price, menuItem.Description, menuItem.RestaurantID).
		Count(&count)

	return count > 0
}
