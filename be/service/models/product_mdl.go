package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       float64
	Images      pq.StringArray `gorm:"type:text[]"`
	Description string
	CategoryID  uint
	Category    Category
}
