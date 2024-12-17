package models

import (
	"github.com/harsh/go-bookstore/pkg/config"
	"github.com/harsh/go-bookstore/pkg/table"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// Initialize the database connection and migrate the schema
func init() {
	config.Connect()
	db = config.GetDB()
	// AutoMigrate automatically migrates the schema to match the models
	if err := db.AutoMigrate(&table.Book{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	} else {
		log.Println("Database schema migrated successfully")
	}
}

// CreateBook creates a new book record in the database
func CreateBook(b *table.Book) *table.Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database
func GetAllBooks() []table.Book {
	var books []table.Book
	db.Find(&books)
	return books
}

// GetBookById retrieves a book record by its ID
func GetBookById(id int64) (*table.Book, *gorm.DB) {
	var book table.Book
	result := db.Where("ID = ?", id).Find(&book)
	return &book, result
}

// DeleteBook deletes a book record by its ID
func DeleteBook(id int64) *table.Book {
	var book table.Book
	db.Where("ID = ?", id).Delete(&book)
	return &book
}
