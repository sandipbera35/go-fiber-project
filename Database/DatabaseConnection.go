package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func connectToPostgreSQL() (*gorm.DB, error) {
	dsn := "user=postgres password=1213 dbname=postgres host=localhost port=5436 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

func Migrator() {
	db, err := connectToPostgreSQL()
	if err != nil {
		panic("Error in database connections")
	}
	Connection = db
	log.Println("Database Connected Successfully...")
	fmt.Println("")
	// db.AutoMigrate()
	// _= db
}
