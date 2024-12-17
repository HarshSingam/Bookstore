package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Connect initializes the database connection
func Connect() {
	dsn := "root:Harsh@7777@tcp(127.0.0.1:3306)/databasefirst?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the database connection
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Database connected successfully!")
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
