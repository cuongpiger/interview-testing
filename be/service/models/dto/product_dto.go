package dto

type (
	ListProductsForm struct {
		Page  int `form:"page"`
		Limit int `form:"limit"`
	}
)
