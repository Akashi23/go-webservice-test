package main

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// If we need to connect to another database, we can add abstraction with interface
//
//	type Database interface {
//		Connect()
//	}
//
// Then we can create a struct that implements this interface
// Here I don't use this approach because I don't need another abstraction
func ConnectDatabase() {
	var database *gorm.DB
	var err error

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(sqlite.Open("inner.db"), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}
	database.AutoMigrate(&Citizen{})

	DB = database
}
