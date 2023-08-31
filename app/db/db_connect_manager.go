package db

import (
	"fmt"
	"log"
	"os"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

	if err != nil {
		log.Fatal("Failed to connect to database")
		os.Exit(2)
	}

	log.Println("Connected to database on port $s", os.Getenv("DB_PORT"))
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = DbInstance{
		Db: db,
	}
}