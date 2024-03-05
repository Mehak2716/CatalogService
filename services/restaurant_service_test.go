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

func setUpRestaurantServiceTest() (sqlmock.Sqlmock, *RestaurantService) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := repository.RestaurantRepository{DB: gormDB}

	service := RestaurantService{repo}
	return mock, &service
}

func TestCreateDuplicateRestaurantExpectAlreadyExistError(t *testing.T) {
	mock, service := setUpRestaurantServiceTest()
	req := &pb.CreateRestaurantRequest{
		Name:     "testName",
		Location: &pb.Location{},
		Status:   "open",
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

func TestCreateRestaurantSuccessfully(t *testing.T) {
	mock, service := setUpRestaurantServiceTest()
	req := &pb.CreateRestaurantRequest{
		Name:     "testName",
		Location: &pb.Location{},
		Status:   "open",
	}

	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rowsCount)

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "status"}).AddRow(1, "testName", "Open")
	mock.ExpectQuery("INSERT INTO \"restaurants\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := service.Create(req)

	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got %", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFetchAllRestaurantsSuccessfully(t *testing.T) {
	mock, service := setUpRestaurantServiceTest()
	req := &pb.FetchAllRestaurantsRequest{}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(1, "Restaurant1", "Open").
		AddRow(2, "Restaurant2", "Closed")
	mock.ExpectQuery("SELECT \\* FROM \"restaurants\"").WillReturnRows(rows)
	res, err := service.FetchAll(req)

	if res == nil {
		t.Fatalf("Expected response but got nil")
	}
	if err != nil {
		t.Fatal("Expected error to be nil but got %", err.Error())
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
