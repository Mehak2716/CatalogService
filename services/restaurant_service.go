package services

import (
	"catalog/mapper"
	pb "catalog/proto/catalog"
	"catalog/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RestaurantService struct {
	Repo repository.RestaurantRepository
}

func (service *RestaurantService) Create(req *pb.CreateRestaurantRequest) (*pb.RestaurantResponse, error) {

	restaurant := mapper.MapToRestaurant(req)
	if service.Repo.IsExists(restaurant) {
		return nil, status.Errorf(codes.AlreadyExists, "Restaurant with the same name already exists")
	}

	createdRestaurant, err := service.Repo.Save(&restaurant)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create restaurant")
	}

	response := mapper.MapToRestaurantResponse(*createdRestaurant)
	return response, nil
}

func (service *RestaurantService) FetchAll(req *pb.FetchAllRestaurantsRequest) (*pb.RestaurantsResponse, error) {

	restaurants, err := service.Repo.FetchAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get restaurants")
	}

	response := mapper.MapToResaurantsResponse(restaurants)
	return response, nil

}
