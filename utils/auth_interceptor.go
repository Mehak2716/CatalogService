package utils

import (
	"context"
	"log"

	pb "catalog/proto/catalog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	switch info.FullMethod {
	case "/Catalog/CreateRestaurant":

		if err := isAdmin(ctx); err != nil {
			return nil, err
		}
		r := req.(*pb.CreateRestaurantRequest)
		if r.Name == "" || r.Location == nil || r.Location.XCordinate == 0 || r.Location.YCordinate == 0 {
			return nil, status.Error(codes.InvalidArgument, "To create restaurant , name and location is required.")
		}

	case "/Catalog/CreateMenuItem":

		if err := isAdmin(ctx); err != nil {
			return nil, err
		}
		r := req.(*pb.CreateMenuItemRequest)
		if r.Name == "" || r.Description == "" || r.RestaurantID == 0 || r.Price == 0 {
			return nil, status.Error(codes.InvalidArgument, "Provide proper details for creating menu item ")
		}
	case "/Catalog/FetchRestaurantMenuItems":

		r := req.(*pb.FetchMenuItemsRequest)
		if r.RestaurantID == 0 {
			return nil, status.Error(codes.InvalidArgument, "Provide restaurant id.")
		}
	}
	resp, err := handler(ctx, req)
	return resp, err
}
