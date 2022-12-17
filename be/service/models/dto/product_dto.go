package dto

import (
	"app/service/models"
	"github.com/lib/pq"
	"strings"
)

type (
	ListProductsForm struct {
		Page  int    `form:"page"`
		Limit int    `form:"limit"`
		Order string `form:"order"`
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

// ListProductsForm's collection of methods
func (s *ListProductsForm) GetOrder() [][]string {
	orders := make([][]string, 0)
	if s.Order == "" {
		return orders
	}

	for _, rule := range strings.Split(strings.ToLower(s.Order), "|") {
		order := strings.Split(rule, ":")
		if len(order) != 2 || (order[1] != "asc" && order[1] != "desc") {
			continue
		}

		orders = append(orders, order)
	}

	return orders
}

// ListProductsResponse's collection of methods
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
