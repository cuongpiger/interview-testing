package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID     int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement;"`
	Name   string `json:"name" gorm:"column:name;not null;type:string;"`
	Parent int64  `json:"parent" gorm:"column:parent;type:int;"`
}
