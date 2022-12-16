package http

import "github.com/gin-gonic/gin"

func (s *ProductHandler) ProductAPIRoute(router *gin.RouterGroup) {
	router.GET("", s.listProducts())
}
