package http

import "github.com/gin-gonic/gin"

func (s *productHandler) ProductAPIRoute(router *gin.RouterGroup) {
	router.GET("/products", s.listProducts())
}
