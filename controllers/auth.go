package controllers

import (
	"net/http"
	"sushi-backend/controllers/dependencies"
	"sushi-backend/services/interfaces"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type AuthController struct {
	authService interfaces.IAuthService
}

func NewAuthController(deps dependencies.AuthControllerDependencies) *AuthController {
	return &AuthController{
		authService: deps.AuthService,
	}
}

func (controller *AuthController) Authorize(_ http.ResponseWriter, r *http.Request) *responses.Response {
	passwordInBase64String := r.Header.Get("X-Admin-Password")

	if passwordInBase64String == "" {
		return responses.NewBadRequestResponse("X-Admin-Password header is required")
	}

	return controller.authService.Authorize(r.Context(), utils.GetClientIpFromContext(r.Context()), passwordInBase64String)
}

func (controller *AuthController) Verify(_ http.ResponseWriter, r *http.Request) *responses.Response {
	token := r.Header.Get("Authorization")

	if token == "" {
		return responses.NewBadRequestResponse("Authorization header is required")
	}

	return controller.authService.Verify(r.Context(), utils.GetClientIpFromContext(r.Context()), token)
}

func (controller *AuthController) Refresh(_ http.ResponseWriter, r *http.Request) *responses.Response {
	token := r.Header.Get("Authorization")

	if token == "" {
		return responses.NewBadRequestResponse("Authorization header is required")
	}

	return controller.authService.Refresh(r.Context(), utils.GetClientIpFromContext(r.Context()), token)
}
