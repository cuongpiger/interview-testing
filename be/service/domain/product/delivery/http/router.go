package http

import "github.com/gin-gonic/gin"

func (s *ProductHandler) ProductAPIRoute(router *gin.RouterGroup) {
	router.GET("/products", s.listProducts())
	router.GET("/categories", s.getAllCategories())
}
