package controllers

import (
	"net/http"
	"sushi-backend/config"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type ProductImageController struct {
	config              config.IConfig
	logger              logger.ILogger
	productImageService interfaces.IProductImageService
}

func NewProductImageController(deps dependencies.ProductImageControllerDependencies) *ProductImageController {
	return &ProductImageController{
		config:              deps.Config,
		logger:              deps.Logger,
		productImageService: deps.ProductImageService,
	}
}

func (h *ProductImageController) Create(_ http.ResponseWriter, r *http.Request) *responses.Response {
	if err := r.ParseMultipartForm(h.config.MaxFileSizeInMb() << 20); err != nil {
		return responses.NewContentTooLargeResponse()
	}

	id, err := utils.GetUUIDIdParam(r)
	if err != nil {
		return err
	}

	file, handler, formFileErr := r.FormFile("file")
	if formFileErr != nil {
		return responses.NewBadRequestResponse("form file is required")
	}

	return h.productImageService.Create(id, file, handler)
}

func (h *ProductImageController) GetById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUUIDIdParam(r)

	if err != nil {
		return err
	}

	return h.productImageService.GetById(id)
}

func (h *ProductImageController) DeleteById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetUUIDIdParam(r)

	if err != nil {
		return err
	}

	return h.productImageService.DeleteById(id)
}
