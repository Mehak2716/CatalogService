package mapper

import (
	"catalog/models"
	pb "catalog/proto/catalog"
)

func MapToMenuItem(req *pb.CreateMenuItemRequest) models.MenuItem {

	return models.MenuItem{
		Name:         req.Name,
		Price:        req.Price,
		Description:  req.Description,
		RestaurantID: req.RestaurantID,
	}
}

func MapToMenuItemResponse(menuItem models.MenuItem) *pb.MenuItemResponse {
	return &pb.MenuItemResponse{
		RestaurantID: int64(menuItem.RestaurantID),
		MenuItem: &pb.MenuItem{
			ID:          menuItem.ID,
			Name:        menuItem.Name,
			Description: menuItem.Description,
			Price:       menuItem.Price,
		},
	}
}

func MapToMenuItemsResponse(restaurantID int64, menuItems []models.MenuItem) *pb.MenuItemsResponse {
	var items []*pb.MenuItem
	for _, menuItem := range menuItems {
		item := &pb.MenuItem{
			ID:          menuItem.ID,
			Name:        menuItem.Name,
			Description: menuItem.Description,
			Price:       menuItem.Price,
		}
		items = append(items, item)
	}

	return &pb.MenuItemsResponse{
		RestaurantID: restaurantID,
		MenuItems:    items,
	}

}
