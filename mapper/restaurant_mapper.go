package mapper

import (
	"catalog/models"
	pb "catalog/proto/catalog"
)

func MapToRestaurant(req *pb.CreateRestaurantRequest) models.Restaurant {

	return models.Restaurant{
		Name: req.Name,
		Location: models.Location{
			XCordinate: float64(req.Location.XCordinate),
			YCordinate: float64(req.Location.YCordinate),
		},
		Status: req.Status,
	}
}

func MapToRestaurantResponse(restaurant models.Restaurant) *pb.RestaurantResponse {

	return &pb.RestaurantResponse{
		ID:   int64(restaurant.ID),
		Name: restaurant.Name,
		Location: &pb.Location{
			XCordinate: float32(restaurant.Location.XCordinate),
			YCordinate: float32(restaurant.Location.YCordinate),
		},
		Status: restaurant.Status,
	}
}

func MapToResaurantsResponse(restaurants []models.Restaurant) *pb.RestaurantsResponse {
	var restaurantResponses []*pb.RestaurantResponse
	for _, restaurant := range restaurants {
		restaurantResponse := &pb.RestaurantResponse{
			ID:   int64(restaurant.ID),
			Name: restaurant.Name,
			Location: &pb.Location{
				XCordinate: float32(restaurant.Location.XCordinate),
				YCordinate: float32(restaurant.Location.YCordinate),
			},
			Status: restaurant.Status,
		}
		restaurantResponses = append(restaurantResponses, restaurantResponse)
	}

	return &pb.RestaurantsResponse{
		Restaurants: restaurantResponses,
	}
}
