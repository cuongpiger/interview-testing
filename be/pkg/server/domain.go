package server

import (
	productUC "app/service/domain/product/usecase"
)

type Domains struct {
	product productUC.IProductUsecase
}
