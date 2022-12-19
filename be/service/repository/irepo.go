package repository

import (
	"app/pkg/config"
	"app/service/repository/postgres"
	"context"
	"gorm.io/gorm"
	"sync"
)

var (
	// product
	productOnce sync.Once
	productRepo postgres.IProductRepo

	// category
	categoryOnce sync.Once
	categoryRepo postgres.ICategoryRepo
)

type IRepo interface {
	NewPostgresProduct() postgres.IProductRepo
	NewPostgresCategory() postgres.ICategoryRepo
}

type repo struct {
	ctx      context.Context
	cfg      *config.AppConfig
	postgres *gorm.DB
}

func NewRepo(ctx context.Context, cfg *config.AppConfig, postgres *gorm.DB) IRepo {
	return &repo{
		ctx:      ctx,
		cfg:      cfg,
		postgres: postgres,
	}
}

func (s *repo) NewPostgresProduct() postgres.IProductRepo {
	productOnce.Do(func() {
		productRepo = postgres.NewProductRepo(s.postgres)
	})

	return productRepo
}

func (s *repo) NewPostgresCategory() postgres.ICategoryRepo {
	categoryOnce.Do(func() {
		categoryRepo = postgres.NewCategoryRepo(s.postgres)
	})

	return categoryRepo
}