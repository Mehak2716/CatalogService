package services

import (
	"catalog/mapper"
	"catalog/repository"

	pb "catalog/proto/catalog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MenuItemService struct {
	Repo repository.MenuItemRepository
}

func (service *MenuItemService) Create(req *pb.CreateMenuItemRequest) (*pb.MenuItemResponse, error) {

	menuItem := mapper.MapToMenuItem(req)
	if service.Repo.IsExists(menuItem) {
		return nil, status.Errorf(codes.AlreadyExists, "Same menu item already exists")
	}
	createdMenuItem, err := service.Repo.Save(&menuItem)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create menu item")
	}
	response := mapper.MapToMenuItemResponse(*createdMenuItem)
	return response, nil
}

func (service *MenuItemService) FetchAll(req *pb.FetchMenuItemsRequest) (*pb.MenuItemsResponse, error) {

	restaurantId := req.RestaurantID

	menuItems, err := service.Repo.FetchAll(int(restaurantId))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get menu items")
	}
	response := mapper.MapToMenuItemsResponse(restaurantId, menuItems)
	return response, nil
}
