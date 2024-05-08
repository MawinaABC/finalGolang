package initializers

import (
	"fmt"
	"github.com/MawinaABC/finalGolang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Order{})
}
