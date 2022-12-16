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

	postgresDB := s.initDatabase()

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

func (s *Server) initDomain() *Domains {
	return &Domains{
		product:
	}
}
