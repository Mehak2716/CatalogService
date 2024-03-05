package services

import (
	pb "catalog/proto/catalog"
	"catalog/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpMenuItemServiceTest() (sqlmock.Sqlmock, *MenuItemService) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := repository.MenuItemRepository{DB: gormDB}

	service := MenuItemService{repo}
	return mock, &service
}

func TestCreateDuplicateMenuItemExpectAlreadyExistError(t *testing.T) {
	mock, service := setUpMenuItemServiceTest()
	req := &pb.CreateMenuItemRequest{
		Name:         "Burger",
		Price:        9.99,
		Description:  "Delicious Burger",
		RestaurantID: 1,
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	res, err := service.Create(req)

	if res != nil {
		t.Fatalf("Expected response to be nil")
	}
	if err != nil {
		gRPCStatus, ok := status.FromError(err)
		if !ok {
			t.Fatal("Expected gRPC status error but got a different type of error")
		}
		expectedStatusCode := codes.AlreadyExists
		if gRPCStatus.Code() != expectedStatusCode {
			t.Fatalf("Expected error code: %v, but got: %v", expectedStatusCode, gRPCStatus.Code())
		}
	} else {
		t.Fatal("Expected an error, but got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreateMenuItemSuccessfully(t *testing.T) {
	mock, service := setUpMenuItemServiceTest()
	req := &pb.CreateMenuItemRequest{
		Name:         "Pizza",
		Price:        12.99,
		Description:  "Tasty Pizza",
		RestaurantID: 1,
	}

	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rowsCount)
	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "restaurant_id"}).
		AddRow(1, "Pizza", 12.99, "Tasty Pizza", 1)
	mock.ExpectQuery("INSERT INTO \"menu_items\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := service.Create(req)

	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchAllMenuItemsSuccessfully(t *testing.T) {
	mock, service := setUpMenuItemServiceTest()
	req := &pb.FetchMenuItemsRequest{
		RestaurantID: 1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "price", "description", "restaurant_id"}).
		AddRow(1, "Burger", 9.99, "Delicious Burger", 1).
		AddRow(2, "Pizza", 12.99, "Tasty Pizza", 1)
	mock.ExpectQuery("SELECT \\* FROM \"menu_items\"").WillReturnRows(rows)
	res, err := service.FetchAll(req)

	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
