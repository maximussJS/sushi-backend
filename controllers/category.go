package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/pkg/logger"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type CategoryController struct {
	logger          logger.ILogger
	categoryService interfaces.ICategoryService
}

func NewCategoryController(deps dependencies.CategoryControllerDependencies) *CategoryController {
	return &CategoryController{
		logger:          deps.Logger,
		categoryService: deps.CategoryService,
	}
}

func (h *CategoryController) Create(w http.ResponseWriter, r *http.Request) *responses.Response {
	var req requests.CreateCategoryRequest

	err := utils.DecodeJSONBody(w, r, &req)
	if err != nil {
		return err
	}

	return h.categoryService.Create(req)
}

func (h *CategoryController) GetAll(_ http.ResponseWriter, r *http.Request) *responses.Response {
	limit, err := utils.GetLimitQueryParam(r, 100)

	if err != nil {
		return err
	}

	offset, err := utils.GetOffsetQueryParam(r, 0)

	if err != nil {
		return err
	}

	return h.categoryService.GetAll(limit, offset)
}

func (h *CategoryController) GetById(_ http.ResponseWriter, r *http.Request) *responses.Response {
	id, err := utils.GetIdParam(r)

	if err != nil {
		return err
	}

	return h.categoryService.GetById(id)
}