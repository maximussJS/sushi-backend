package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/pkg/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
)

type ProductController struct {
	logger         logger.ILogger
	ProductService interfaces.IProductService
}

func NewProductController(deps dependencies.ProductControllerDependencies) *ProductController {
	return &ProductController{
		logger:         deps.Logger,
		ProductService: deps.ProductService,
	}
}

func (h *ProductController) GetAll(w http.ResponseWriter, r *http.Request) *responses.Response {
	return h.ProductService.GetAll()
}
