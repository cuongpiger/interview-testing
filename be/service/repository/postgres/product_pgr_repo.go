package postgres

import (
	"app/service/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IProductRepo interface {
	ListProducts(offset, limit int) ([]models.Product, error)
}

type productRepo struct {
	client *gorm.DB
	log    *zap.SugaredLogger
}

func NewProductRepo(client *gorm.DB) IProductRepo {
	return &productRepo{
		client: client,
		log:    zap.S(),
	}
}

func (s *productRepo) ListProducts(offset, limit int) ([]models.Product, error) {
	var products []models.Product
	err := s.client.Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		s.log.Errorf("[repo][product] ListProducts: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	return products, nil
}
