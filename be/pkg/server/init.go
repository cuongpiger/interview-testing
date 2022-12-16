package server

import (
	"app/service/domain/product/delivery/http"
	productUC "app/service/domain/product/usecase"
	"app/service/models"
	"app/service/repository"
	"github.com/gin-contrib/cors"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Domains struct {
	product productUC.IProductUsecase
}

func (s *Server) initCORS() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"*",
		"Origin",
		"Content-Length",
		"Content-Type",
		"Authorization",
		"X-Access-Token",
	}
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) initDomains(repo repository.IRepo) *Domains {
	return &Domains{
		product: productUC.NewProductUsecase(s.cfg, repo),
	}
}

func (s *Server) initRouter(domains *Domains) {
	productHdl := http.NewProductHandler(s.cfg, domains.product)
	productHdl.ProductAPIRoute(s.router.Group("/api/v1/products"))
}

func (s *Server) initDatabase() *gorm.DB {
	postgresDB, errDB := gorm.Open(postgres.Open(s.cfg.PostgresDSN), &gorm.Config{})
	if errDB != nil {
		zap.S().Errorf("Failed to connect to postgres: %v", errDB)
		panic(errDB)
	}

	errDB = postgresDB.AutoMigrate(&models.Category{})
	errDB = postgresDB.AutoMigrate(&models.Product{})

	if errDB != nil {
		zap.S().Errorf("Failed to migrate postgres: %v", errDB)
		panic(errDB)
	}

	return postgresDB
}
