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

	clientIP := utils.GetClientIpFromContext(r.Context())

	return controller.authService.Authorize(clientIP, passwordInBase64String)
}

func (controller *AuthController) Verify(_ http.ResponseWriter, r *http.Request) *responses.Response {
	token := r.Header.Get("Authorization")

	if token == "" {
		return responses.NewBadRequestResponse("Authorization header is required")
	}

	clientIP := utils.GetClientIpFromContext(r.Context())

	return controller.authService.Verify(clientIP, token)
}
