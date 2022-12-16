package postgres

import (
	"gorm.io/gorm"
)

type ICategoryRepo interface {
}

type categoryRepo struct {
	client *gorm.DB
}

func NewCategoryRepo(client *gorm.DB) ICategoryRepo {
	return &categoryRepo{
		client: client,
	}
}
