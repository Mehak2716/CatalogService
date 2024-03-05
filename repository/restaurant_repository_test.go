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
	expectedRestaurants := []models.Restaurant{
		{Name: "Restaurant1", Status: "Open"},
		{Name: "Restaurant2", Status: "Closed"},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status"}).
		AddRow(1, "Restaurant1", "Open").
		AddRow(2, "Restaurant2", "Closed")
	mock.ExpectQuery("SELECT \\* FROM \"restaurants\"").WillReturnRows(rows)
	fetchedRestaurants, err := restaurantRepo.FetchAll()

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if len(fetchedRestaurants) != len(expectedRestaurants) {
		t.Fatalf("Unexpected number of restaurants. Expected: %d, Got: %d", len(expectedRestaurants), len(fetchedRestaurants))
	}
	for i, expected := range expectedRestaurants {
		if fetchedRestaurants[i].Name != expected.Name || fetchedRestaurants[i].Status != expected.Status {
			t.Fatalf("Unexpected restaurant. Expected: %+v, Got: %+v", expected, fetchedRestaurants[i])
		}
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
