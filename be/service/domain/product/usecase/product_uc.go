package usecase

import (
	"app/pkg/config"
	"app/pkg/request"
	"app/service/models/dto"
	"app/service/repository"
	"go.uber.org/zap"
)

type IProductUsecase interface {
	ListProducts(page, limit int, filter []*request.Filter, orders [][]string) ([]*dto.ListProductsResponse, error)
	GetAllCategories() ([]*dto.GetAllCategoriesResponse, error)
}

type productUsecase struct {
	cfg  *config.AppConfig
	log  *zap.SugaredLogger
	repo repository.IRepo
}

func NewProductUsecase(cfg *config.AppConfig, repo repository.IRepo) IProductUsecase {
	return &productUsecase{
		cfg:  cfg,
		log:  zap.S(),
		repo: repo,
	}
}

func (s *productUsecase) ListProducts(page, limit int, filter []*request.Filter, orders [][]string) ([]*dto.ListProductsResponse, error) {
	offset := (page - 1) * limit
	products, err := s.repo.NewPostgresProduct().ListProducts(offset, limit, filter, orders)
	if err != nil {
		s.log.Errorf("[usecase][product] ListProducts: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	res := make([]*dto.ListProductsResponse, len(products))
	for i, product := range products {
		res[i] = new(dto.ListProductsResponse).Merge(&product)
	}

	return res, nil
}

func (s *productUsecase) GetAllCategories() ([]*dto.GetAllCategoriesResponse, error) {
	categories, err := s.repo.NewPostgresCategory().GetAllCategories()
	if err != nil {
		s.log.Errorf("[usecase][product] GetAllCategories: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	res := make([]*dto.GetAllCategoriesResponse, len(categories))
	for i, category := range categories {
		res[i] = new(dto.GetAllCategoriesResponse).Merge(&category)
	}

	return res, nil
}
