package utils

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/raulaguila/api-go/api/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	strConnection := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now()
		},
	})

	if err != nil {
		log.Fatalf("Error connecting database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&domain.Profile{})
	db.AutoMigrate(&domain.User{})
	return db
}
