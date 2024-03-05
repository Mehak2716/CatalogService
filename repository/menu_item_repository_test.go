package repository

import (
	"catalog/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpMenuItemTest() (sqlmock.Sqlmock, *MenuItemRepository) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	repo := MenuItemRepository{DB: gormDB}

	return mock, &repo
}
func TestMenuItemCreatedSuccessfully(t *testing.T) {

	mock, menuItemRepo := setUpMenuItemTest()
	menuItem := models.MenuItem{

		Name:        "testName",
		Price:       10,
		Description: "testing",
	}

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "testName", 10)
	mock.ExpectQuery("INSERT INTO \"menu_items\"").WillReturnRows(rows)
	mock.ExpectCommit()
	res, err := menuItemRepo.Save(&menuItem)

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

func TestFetchAllMenuItemsSuccessfully(t *testing.T) {
	mock, menuItemRepo := setUpMenuItemTest()
	expectedMenuItems := []models.MenuItem{
		{Name: "Item1", Price: 10, Description: "Description1"},
		{Name: "Item2", Price: 15, Description: "Description2"},
	}
	restaurantID := 1

	rows := sqlmock.NewRows([]string{"id", "name", "price", "description"}).
		AddRow(1, "Item1", 10, "Description1").
		AddRow(2, "Item2", 15, "Description2")
	mock.ExpectQuery("SELECT \\* FROM \"menu_items\" WHERE restaurant_id = ?").WithArgs(restaurantID).WillReturnRows(rows)
	fetchedMenuItems, err := menuItemRepo.FetchAll(restaurantID)

	if err != nil {
		t.Fatalf("Error not expected but encountered: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if len(fetchedMenuItems) != len(expectedMenuItems) {
		t.Fatalf("Unexpected number of menu items. Expected: %d, Got: %d", len(expectedMenuItems), len(fetchedMenuItems))
	}
	for i, expected := range expectedMenuItems {
		if fetchedMenuItems[i].Name != expected.Name || fetchedMenuItems[i].Price != expected.Price || fetchedMenuItems[i].Description != expected.Description {
			t.Fatalf("Unexpected menu item. Expected: %+v, Got: %+v", expected, fetchedMenuItems[i])
		}
	}
}

func TestMenuItemExistsSuccessfully(t *testing.T) {
	mock, menuItemRepo := setUpMenuItemTest()
	menuItem := models.MenuItem{
		Name:         "Item1",
		Price:        10,
		Description:  "Description1",
		RestaurantID: 1,
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	exists := menuItemRepo.IsExists(menuItem)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if !exists {
		t.Fatalf("Expected menu item to exist, but it does not")
	}
}

func TestMenuItemDoesNotExistsSuccessfully(t *testing.T) {
	mock, menuItemRepo := setUpMenuItemTest()
	menuItem := models.MenuItem{
		Name:         "Item1",
		Price:        10,
		Description:  "Description1",
		RestaurantID: 1,
	}

	rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
	mock.ExpectQuery("SELECT count(.+) FROM (.+)").
		WillReturnRows(rows)
	exists := menuItemRepo.IsExists(menuItem)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	if exists {
		t.Fatalf("Expected menu item to not exist, but it exist")
	}
}
