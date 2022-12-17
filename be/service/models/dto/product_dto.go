package dto

import (
	"app/service/models"
	"github.com/lib/pq"
)

type (
	ListProductsForm struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}

	ListProductsResponse struct {
		ID          int64          `json:"id"`
		Name        string         `json:"name"`
		Price       float64        `json:"price"`
		Images      pq.StringArray `json:"images"`
		Description string         `json:"description"`
		Category    int64          `json:"category"`
	}
)

func (s *ListProductsResponse) Merge(data interface{}) *ListProductsResponse {
	switch d := data.(type) {
	case *models.Product:
		return s.mergeProduct(d)
	}

	return s
}

func (s *ListProductsResponse) mergeProduct(product *models.Product) *ListProductsResponse {
	s.ID = product.ID
	s.Name = product.Name
	s.Price = product.Price
	s.Images = product.Images
	s.Description = product.Description
	s.Category = product.Category
	return s
}
