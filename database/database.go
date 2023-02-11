package database

import (
	"fmt"
	"log"

	"github.com/zakirkun/grpc-crud/app/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	host := "localhost"
	port := "5432"
	dbName := "movie"
	dbUser := "nakano"
	password := "Nakano@Miku!"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
		return nil, err
	}

	fmt.Println("Database connection successful...")

	db.AutoMigrate(&model.Movie{})

	return db, nil
}
