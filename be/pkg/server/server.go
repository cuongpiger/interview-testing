package server

import (
	"app/service/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	cfg    *AppConfig
}

func NewServer(cfg *AppConfig) (*Server, error) {
	router := gin.Default()
	return &Server{
		router: router,
		cfg:    cfg,
	}, nil
}

func (s *Server) Init() {
	//ctx := context.Background()

	postgresDB, errDB := gorm.Open(postgres.Open(s.cfg.PostgresDSN), &gorm.Config{})
	if errDB != nil {
		zap.S().Errorf("Failed to connect to postgres: %v", errDB)
		panic(errDB)
	}

	_ = postgresDB.AutoMigrate(&models.Product{})

	var product models.Product
	postgresDB.First(&product)

	fmt.Println(product)
}
