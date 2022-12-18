package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey,autoIncrement"`
	Name     string `gorm:"unique"`
	ParentID uint   `gorm:"default:null"`
	Parent   *Category
}
