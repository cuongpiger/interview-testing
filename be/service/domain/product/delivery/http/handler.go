package http

import (
	"app/pkg/config"
	"app/pkg/request"
	productUC "app/service/domain/product/usecase"
	"app/service/models/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	cfg       *config.AppConfig
	productUC productUC.IProductUsecase
}

func NewProductHandler(cfg *config.AppConfig, productUC productUC.IProductUsecase) *ProductHandler {
	return &ProductHandler{
		cfg:       cfg,
		productUC: productUC,
	}
}

func (s *ProductHandler) listProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form := new(dto.ListProductsForm)
		err := ctx.BindQuery(form)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest,
				request.NewResponse().
					SetCode("BAD_REQUEST").
					SetMessage(fmt.Sprintf("%+v", err)).GetResponse())
			return
		}

		products, err := s.productUC.ListProducts(form.Page, form.Limit, form.GetFilter(), form.GetOrder())
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest,
				request.NewResponse().
					SetCode("BAD_REQUEST").
					SetMessage(fmt.Sprintf("%+v", err)).GetResponse())
			return
		}

		ctx.PureJSON(http.StatusOK, request.NewResponse().SetCode("SUCCESS").SetData(products).GetResponse())
	}
}
