package postgres

import (
	"app/pkg/request"
	"app/service/models"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IProductRepo interface {
	ListProducts(offset, limit int, filter []*request.Filter, orders [][]string) ([]models.Product, error)
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

func (s *productRepo) ListProducts(offset, limit int, filter []*request.Filter, orders [][]string) ([]models.Product, error) {
	var (
		products []models.Product
		query    = s.getTable().Offset(offset).Limit(limit).Preload("Category")
	)

	// filter phase
	if len(filter) > 0 {
		for _, f := range filter {
			if !s.validCommand(f.Field, f.Operator) {
				continue
			}

			switch f.Operator {
			case "eq":
				query = query.Where(f.GetQuery(dbName), f.Value)
			case "like":
				query = query.Where(f.GetQuery(dbName), "%"+f.Value.(string)+"%")
			case "in":
				query = query.Where(f.GetQuery(dbName), f.Value.([]int))
			case "gte", "lte":
				switch f.Field {
				case "price":
					query = query.Where(f.GetQuery(dbName), cast.ToFloat64(f.Value))
				case "category_id", "id":
					query = query.Where(f.GetQuery(dbName), cast.ToInt(f.Value))
				}
			}
		}
	}

	// sort phase
	if len(orders) > 0 {
		for _, order := range orders {
			query = query.Order(order[0] + " " + order[1])
		}
	}

	query = query.Find(&products)
	if err := query.Error; err != nil {
		s.log.Errorf("[repo][product] ListProducts: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	return products, nil
}

func (s *productRepo) validCommand(field, operation string) bool {
	switch field {
	case "id":
		switch operation {
		case "eq", "in", "gte", "lte":
			return true
		}
	case "name":
		switch operation {
		case "eq", "like":
			return true
		}
	case "price":
		switch operation {
		case "eq", "gte", "lte":
			return true
		}
	case "category_id":
		switch operation {
		case "eq", "in", "gte", "lte":
			return true
		}
	case "description":
		switch operation {
		case "eq", "like":
			return true
		}
	}

	return false
}
