package services

import (
	"encoding/base64"
	"sushi-backend/config"
	"sushi-backend/internal/jwt"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
)

type AuthService struct {
	config     config.IConfig
	jwtService jwt.IJwtService
}

func NewAuthService(deps dependencies.AuthServiceDependencies) *AuthService {
	return &AuthService{
		config:     deps.Config,
		jwtService: deps.JwtService,
	}
}

func (service *AuthService) Authorize(clientIp, passwordInBase64String string) *responses.Response {
	passwordString, err := base64.StdEncoding.DecodeString(passwordInBase64String)
	if err != nil {
		return responses.NewBadRequestResponse("Invalid password. Should be base64 encoded")
	}

	if string(passwordString) != service.config.AdminPassword() {
		return responses.NewUnauthorizedResponse("Invalid password")
	}

	token := service.jwtService.GenerateTokenWithClientIp(clientIp)

	return responses.NewSuccessResponse(token)
}

func (service *AuthService) Verify(clientIp, token string) *responses.Response {
	clientIpFromJwt, err := service.jwtService.VerifyTokenWithClientIp(token)
	if err != nil {
		return responses.NewUnauthorizedResponse(err.Error())
	}

	if clientIp != clientIpFromJwt {
		return responses.NewUnauthorizedResponse("Invalid token")
	}

	return responses.NewSuccessResponse(nil)
}
