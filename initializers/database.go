package initializers

import (
	"log"
	"os"

	"github.com/Richard-Owen-Tangrady/richard/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	env := godotenv.Load()

	if env != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DB")
	// fmt.Printf("%v", dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database in init function: %v", err)
	}
	log.Println("Database connected successfully via init!")

	if err != nil {
		panic("Failed to connect to database")
	}
}

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
