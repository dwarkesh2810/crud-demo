package initialise

import (
	postmodel "crud/internal/modules/post/postModel"
	"crud/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect DB")
	}

	err = db.AutoMigrate(&postmodel.Posts{}, &models.Users{}, &models.Category{})

	if err != nil {
		log.Fatal("failed to migrate")
	}

	DB = db
}

func Connection() *gorm.DB {
	return DB
}
