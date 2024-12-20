package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type ProductImageController struct {
	logger              logger.ILogger
	productImageService interfaces.IProductImageService
}

func NewProductImageController(deps dependencies.ProductImageControllerDependencies) *ProductImageController {
	return &ProductImageController{
		logger:              deps.Logger,
		productImageService: deps.ProductImageService,
	}
}

func (h *ProductImageController) Create(w http.ResponseWriter, r *http.Request) *responses.Response {
	if err := r.ParseMultipartForm(200 << 20); err != nil {
		return responses.NewContentTooLargeResponse()
	}

	id, err := utils.GetIdParam(r)
	if err != nil {
		return err
	}

	file, handler, error := r.FormFile("file")
	if error != nil {
		return responses.NewBadRequestResponse("form file is required")
	}

	return h.productImageService.Create(id, file, handler)
}

func (h *ProductImageController) GetById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetIdParam(r)

	if err != nil {
		return err
	}

	return h.productImageService.GetById(id)
}

func (h *ProductImageController) DeleteById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetIdParam(r)

	if err != nil {
		return err
	}

	return h.productImageService.DeleteById(id)
}
