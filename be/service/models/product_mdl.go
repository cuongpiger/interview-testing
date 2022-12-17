package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement;"`
	Name        string         `gorm:"column:name;not null;type:string;"`
	Price       float64        `gorm:"column:price;not null;type:float;"`
	Images      pq.StringArray `gorm:"column:images;type:text[];"`
	Description string         `gorm:"column:description;type:text;"`
	Category    int64          `gorm:"column:category;type:int;references:ID;"`
}

type ProductCategory struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement;"`
	Name        string         `gorm:"column:name;not null;type:string;"`
	Price       float64        `gorm:"column:price;not null;type:float;"`
	Images      pq.StringArray `gorm:"column:images;type:text[];"`
	Description string         `gorm:"column:description;type:text;"`
	Category    Category       `gorm:"foreignKey:ID;"`
}
