package database

import (
	"fmt"
	"log"
	"os"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
}

var DB DbInstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=db user=admin password=12345 dbname=avitodb port=5432 sslmode=disable TimeZone=Europe/Moscow"
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = DbInstance{
		DB: db,
	}
}

func GetConnection() (*gorm.DB, error) {
	if DB.DB != nil {
		return DB.DB, nil
	}
	return nil, fmt.Errorf(`db not connected`)
}