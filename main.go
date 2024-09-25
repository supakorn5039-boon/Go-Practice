package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/supakorn5039-boon/practice/controllers"
	"github.com/supakorn5039-boon/practice/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize the database connection
func initDB() *gorm.DB {
	dsn := "host=hostName user=yourUser password=yourPassword dbname=yourDB port=Port sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

func main() {
	app := fiber.New()
	db := initDB()

	// Automatically migrate the Product model
	db.AutoMigrate(&controllers.Product{})
	controllers.DB = db

	seed.SeedProducts(db)

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Define routes
	app.Get("/api/products", controllers.GetProducts)
	app.Get("/api/product/:id", controllers.GetProductById)
	app.Post("/api/product", controllers.CreateProduct)

	// Start server on port 3010
	if err := app.Listen(":3010"); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
