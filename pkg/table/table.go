package table

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        string `gorm:""json:"id"`
	Name      string `gorm:""json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}
