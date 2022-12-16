package usecase

import (
	"app/pkg/config"
	"app/service/models"
	"app/service/repository"
	"go.uber.org/zap"
)

type IProductUsecase interface {
	ListProducts(page, limit int) ([]models.Product, error)
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

func (s *productUsecase) ListProducts(page, limit int) ([]models.Product, error) {
	offset := (page - 1) * limit
	products, err := s.repo.NewPostgresProduct().ListProducts(offset, limit)
	if err != nil {
		s.log.Errorf("[usecase][product] ListProducts: failed to get data from Postgres: %#v", err)
		return nil, err
	}

	return products, nil
}
