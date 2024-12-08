package controller

import (
	"api/model"
	"api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// Usecases
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: usecase,
	}

}

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "mock",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
