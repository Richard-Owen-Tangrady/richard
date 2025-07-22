package initializers

import (
	"os"

	"github.com/Richard-Owen-Tangrady/richard/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
}

func SyncDataBase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Product{})
}
