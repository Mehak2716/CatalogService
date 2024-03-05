package main

import (
	"catalog/config"
	pb "catalog/proto/catalog"
	"catalog/repository"
	"catalog/server"
	"catalog/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

const servePort string = ":9000"

func main() {
	lis, err := net.Listen("tcp", servePort)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	log.Printf("Listening on %s\n", servePort)
	db := config.DatabaseConnection()
	grpcServer := grpc.NewServer()

	menuItemRepo := repository.MenuItemRepository{DB: db}
	menuItemService := services.MenuItemService{Repo: menuItemRepo}

	restaurantRepo := repository.RestaurantRepository{DB: db}
	restaurantService := services.RestaurantService{Repo: restaurantRepo}

	server := server.NewCatalogServer(menuItemService, restaurantService)
	pb.RegisterCatalogServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
