package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type ProductController struct {
	logger         logger.ILogger
	productService interfaces.IProductService
}

func NewProductController(deps dependencies.ProductControllerDependencies) *ProductController {
	return &ProductController{
		logger:         deps.Logger,
		productService: deps.ProductService,
	}
}

func (h *ProductController) Create(w http.ResponseWriter, r *http.Request) *responses.Response {
	var req requests.CreateProductRequest

	err := utils.DecodeJSONBody(w, r, &req)
	if err != nil {
		return err
	}

	return h.productService.Create(r.Context(), req)
}

func (h *ProductController) GetAll(_ http.ResponseWriter, r *http.Request) *responses.Response {
	limit, err := utils.GetLimitQueryParam(r, 100)

	if err != nil {
		return err
	}

	offset, err := utils.GetOffsetQueryParam(r, 0)

	if err != nil {
		return err
	}

	return h.productService.GetAll(r.Context(), limit, offset)
}

func (h *ProductController) GetById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUUIDIdParam(r)

	if err != nil {
		return err
	}

	return h.productService.GetById(r.Context(), id)
}

func (h *ProductController) UpdateById(w http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUUIDIdParam(r)

	if err != nil {
		return err
	}

	var req requests.UpdateProductRequest

	if err := utils.DecodeJSONBody(w, r, &req); err != nil {
		return err
	}

	return h.productService.UpdateById(r.Context(), id, req)
}

func (h *ProductController) DeleteById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUUIDIdParam(r)

	if err != nil {
		return err
	}

	return h.productService.DeleteById(r.Context(), id)
}
