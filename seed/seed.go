package seed

import (
	"log"

	"github.com/supakorn5039-boon/practice/controllers"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) {

	var count int64
	db.Model(&controllers.Product{}).Count(&count)

	if count == 0 {

		products := []controllers.Product{
			{Name: "Laptop", Price: 1099.99},
			{Name: "Smartphone", Price: 799.99},
			{Name: "Headphones", Price: 199.99},
		}

		if err := db.Create(&products).Error; err != nil {
			log.Fatal("Failed to seed products:", err)
		}
		log.Println("Initial products seeded.")
	} else {
		log.Println("Products already exist in the database.")
	}
}
