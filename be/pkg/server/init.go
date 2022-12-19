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
	productHdl.ProductAPIRoute(s.router.Group("/api/v1"))
}

func (s *Server) initDatabase() *gorm.DB {
	postgresDB, errDB := gorm.Open(postgres.Open(s.cfg.PostgresDSN), &gorm.Config{})
	if errDB != nil {
		zap.S().Errorf("Failed to connect to postgres: %v", errDB)
		panic(errDB)
	}

	errDB = postgresDB.AutoMigrate(
		&models.Category{},
		&models.Product{})

	if errDB != nil {
		zap.S().Errorf("Failed to migrate postgres: %v", errDB)
		panic(errDB)
	}

	return postgresDB
}

func (s *Server) initSampleData(db *gorm.DB) {
	categories := []models.Category{
		{Name: "ROOT"},
		{Name: "Fashion", ParentID: 1},
		{Name: "Shirt", ParentID: 2},
		{Name: "Jacket", ParentID: 2},
		{Name: "Dress", ParentID: 2},
		{Name: "Unisex", ParentID: 2},
	}
	db.Create(&categories)

	products := []models.Product{
		{Name: "Love and Thunder Blazer", CategoryID: 4, Price: 100000, Description: "The pink blazer from Love and Thunder.", Images: []string{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F0.jpg?alt=media&token=7b40f3cc-37a2-414a-8550-cc8b2389dda3", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F0%2F1.jpg?alt=media&token=8faa9613-e27d-4193-91ec-615d2d53dfe5"}},
		{Name: "Black Window Dress", CategoryID: 5, Price: 300000, Description: "Black Window wears it to protect the Earth from Thanos.", Images: []string{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F0.jpg?alt=media&token=cfb5ea77-7215-45a1-946c-78f281fdd4b6", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F1%2F1.jpg?alt=media&token=7ed979aa-2992-4d97-a862-c98007312ad6"}},
		{Name: "Sport Suit IronHeart", CategoryID: 5, Price: 250000, Description: "IronHeart x Shuri in Wakanda forever.", Images: []string{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F0.jpg?alt=media&token=9baece8e-5bb2-458f-b454-9308c6afebf2", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F10%2F1.jpg?alt=media&token=fb028a35-6ddf-4c3c-9634-2db90d202756"}},
		{Name: "Rubber suit", CategoryID: 6, Price: 500000, Description: "This dress trades on people's fears to make you special.", Images: []string{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F0.jpg?alt=media&token=909e31c1-38bc-4de1-ad08-4f5b52f9bad1", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F11%2F1.jpg?alt=media&token=9f2a9325-4c9e-4d04-85f2-a455add0b901"}},
		{Name: "The jacket of Tom Holland", CategoryID: 3, Price: 900000, Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.", Images: []string{"https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F0.jpg?alt=media&token=a4379223-24db-4184-89de-dcf4cb089a15", "https://firebasestorage.googleapis.com/v0/b/clover-78661.appspot.com/o/Product%2F12%2F1.jpg?alt=media&token=9d705a38-8bcf-4daa-9043-5e073e4e0d5a"}},
	}
	db.Create(&products)
}
