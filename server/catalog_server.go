package server

import (
	"catalog/proto/catalog"
	"catalog/services"
	"context"
)

type CatalogServer struct {
	menuItemService   services.MenuItemService
	restaurantService services.RestaurantService
	catalog.CatalogServer
}

func NewCatalogServer(menuItemService services.MenuItemService, restaurantService services.RestaurantService) *CatalogServer {
	return &CatalogServer{
		menuItemService:   menuItemService,
		restaurantService: restaurantService,
	}
}

func (server *CatalogServer) CreateRestaurant(ctx context.Context, req *catalog.CreateRestaurantRequest) (*catalog.RestaurantResponse, error) {
	return server.restaurantService.Create(req)
}

func (server *CatalogServer) CreateMenuItem(ctx context.Context, req *catalog.CreateMenuItemRequest) (*catalog.MenuItemResponse, error) {
	return server.menuItemService.Create(req)
}

func (server *CatalogServer) FetchAllRestaurants(ctx context.Context, req *catalog.FetchAllRestaurantsRequest) (*catalog.RestaurantsResponse, error) {
	return server.restaurantService.FetchAll(req)
}

func (server *CatalogServer) FetchRestaurantMenuItems(ctx context.Context, req *catalog.FetchMenuItemsRequest) (*catalog.MenuItemsResponse, error) {
	return server.menuItemService.FetchAll(req)
}
