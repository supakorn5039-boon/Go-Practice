package controllers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	return gormDB, mock
}

func TestGetProducts(t *testing.T) {
	app := fiber.New()

	// Set up the mock database
	db, mock := setupMockDatabase()
	DB = db // Use the mock database

	// Prepare mock data
	mock.ExpectQuery(`SELECT \* FROM "products"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
			AddRow(1, "Laptop", 1099.99).
			AddRow(2, "Smartphone", 799.99).
			AddRow(3, "Headphones", 199.99))

	app.Get("/products", GetProducts)

	// Create a new request
	req := httptest.NewRequest("GET", "/products", nil)
	res, err := app.Test(req)

	// Check for errors
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer res.Body.Close()

	// Assert the response body
	expected := `[{"id":1,"name":"Laptop","price":1099.99},{"id":2,"name":"Smartphone","price":799.99},{"id":3,"name":"Headphones","price":199.99}]`
	assert.JSONEq(t, expected, string(body))

	// Check for any mock expectations
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
