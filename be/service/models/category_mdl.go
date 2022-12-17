package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID     int64  `gorm:"column:id;primaryKey;autoIncrement;"`
	Name   string `gorm:"column:name;not null;type:string;"`
	Parent int64  `gorm:"column:parent;type:int;"`
}
