package postgres

import (
	"app/service/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IProductRepo interface {
	ListProducts(offset, limit int, orders [][]string) ([]models.ProductCategory, error)
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

func (s *productRepo) getTable() *gorm.DB {
	return s.client.Model(&models.Product{})
}

func (s *productRepo) ListProducts(offset, limit int, orders [][]string) ([]models.ProductCategory, error) {
	var (
		products []models.ProductCategory
		query    = s.getTable().Preload("Category").Offset(offset).Limit(limit)
	)

	// sort phase
	if len(orders) > 0 {
		for _, order := range orders {
			query = query.Order(order[0] + " " + order[1])
		}
	}

	err := query.Find(&products).Error
	if err != nil {
		s.log.Errorf("[repo][product] ListProducts: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	return products, nil
}
