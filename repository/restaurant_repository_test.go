package repository

import (
	"catalog/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpRestaurantTest() (sqlmock.Sqlmock, *RestaurantRepository) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := RestaurantRepository{DB: gormDB}

	return mock, &repo
}
func TestRestaurantCreatedSuccessfully(t *testing.T) {

	mock, restaurantRepo := setUpRestaurantTest()
	restaurant := models.Restaurant{

		Name:     "testMame",
		Location: models.Location{},
		Status:   "Open",
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(1, "testName", "Open")
	mock.ExpectQuery("INSERT INTO \"restaurants\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := restaurantRepo.Save(&restaurant)

	if err != nil {
		t.Fatalf("Error not expected but encountered %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
	if res.ID != 1 || res.Name != "testName" {
		t.Fatal("Unexpected Result")
	}
}

func TestFetchAllRestaurantsSuccessfully(t *testing.T) {
	mock, restaurantRepo := setUpRestaurantTest()

	rows := sqlmock.NewRows([]string{"id", "name", "status", "xcordinate", "ycordinate"}).
		AddRow(1, "Restaurant1", "Open", 1.23, 4.56).
		AddRow(2, "Restaurant2", "Closed", 7.89, 10.11)
	mock.ExpectQuery("SELECT \\* FROM \"restaurants\"").WillReturnRows(rows)

	locationRows := sqlmock.NewRows([]string{"id", "xcordinate", "ycordinate", "restaurant_id"}).
		AddRow(1, 1.23, 4.56, 1).
		AddRow(2, 7.89, 10.11, 2)
	mock.ExpectQuery("SELECT \\* FROM \"locations\" WHERE \"locations\".\"restaurant_id\" IN \\(\\$1,\\$2\\)").
		WithArgs(1, 2).
		WillReturnRows(locationRows)

	_, err := restaurantRepo.FetchAll()

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestIsExistsForExistingRestaurantSuccessfully(t *testing.T) {
	mock, repo := setUpRestaurantTest()
	restaurant := models.Restaurant{
		Name: "testName",
		Location: models.Location{
			XCordinate: 30,
			YCordinate: 40,
		},
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(restaurant)

	if !result {
		t.Fatalf("Expected IsExists to return true, but got false")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestIsExistsForNonExistingRestaurantSuccessfully(t *testing.T) {
	mock, repo := setUpRestaurantTest()
	restaurant := models.Restaurant{
		Name: "testName",
		Location: models.Location{
			XCordinate: 30,
			YCordinate: 40,
		},
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	result := repo.IsExists(restaurant)

	if result {
		t.Fatalf("Expected IsExists to return false, but got true")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
