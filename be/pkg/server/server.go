package server

import (
	"app/pkg/config"
	"app/service/repository"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	server *http.Server
	router *gin.Engine
	cfg    *config.AppConfig
}

func NewServer(cfg *config.AppConfig) (*Server, error) {
	router := gin.Default()
	return &Server{
		router: router,
		cfg:    cfg,
	}, nil
}

func (s *Server) Init() {
	ctx := context.Background()
	s.initCORS()
	postgresDB := s.initDatabase()
	//s.initSampleData(postgresDB)
	repo := repository.NewRepo(ctx, s.cfg, postgresDB)
	domains := s.initDomains(repo)
	s.initRouter(domains)
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.cfg.Port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.router,
	}
	zap.S().Infof("Start server at %s", addr)
	return s.server.ListenAndServe()
}
