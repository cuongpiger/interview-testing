package http

import (
	"app/pkg/server"
	productUC "app/service/domain/product/usecase"
	"app/service/models/dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	cfg       *server.AppConfig
	productUC productUC.IProductUsecase
}

func NewProductHandler(cfg *server.AppConfig, productUC productUC.IProductUsecase) *productHandler {
	return &productHandler{
		cfg:       cfg,
		productUC: productUC,
	}
}

func (s *productHandler) listProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form := new(dto.ListProductsForm)
		err := ctx.BindQuery(form)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest,
				server.NewResponse().
					SetCode("BAD_REQUEST").
					SetMessage(fmt.Sprintf("%+v", err)).GetResponse())
			return
		}

		products, err := s.productUC.ListProducts(form.Page, form.Limit)
		if err != nil {
			ctx.PureJSON(http.StatusBadRequest,
				server.NewResponse().
					SetCode("BAD_REQUEST").
					SetMessage(fmt.Sprintf("%+v", err)).GetResponse())
			return
		}

		ctx.PureJSON(http.StatusOK, server.NewResponse().SetCode("SUCCESS").SetData(products).GetResponse())
	}
}
