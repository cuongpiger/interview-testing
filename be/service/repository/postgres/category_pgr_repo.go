package postgres

import (
	"app/service/models"
	"gorm.io/gorm"
)

type ICategoryRepo interface {
	GetAllCategories() ([]models.Category, error)
}

type categoryRepo struct {
	client *gorm.DB
}

func NewCategoryRepo(client *gorm.DB) ICategoryRepo {
	return &categoryRepo{
		client: client,
	}
}

func (s *categoryRepo) getTable() *gorm.DB {
	return s.client.Model(&models.Category{})
}

func (s *categoryRepo) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := s.getTable().Where("name <> 'ROOT'").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
