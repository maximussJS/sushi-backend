package services

import (
	"context"
	"encoding/base64"
	"sushi-backend/config"
	"sushi-backend/internal/jwt"
	"sushi-backend/models"
	"sushi-backend/repositories/interfaces"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/responses"
	"time"
)

type AuthService struct {
	config            config.IConfig
	jwtService        jwt.IJwtService
	sessionRepository interfaces.ISessionRepository
}

func NewAuthService(deps dependencies.AuthServiceDependencies) *AuthService {
	return &AuthService{
		config:            deps.Config,
		jwtService:        deps.JwtService,
		sessionRepository: deps.SessionRepository,
	}
}

func (service *AuthService) Authorize(ctx context.Context, clientIp, passwordInBase64String string) *responses.Response {
	passwordString, err := base64.StdEncoding.DecodeString(passwordInBase64String)
	if err != nil {
		return responses.NewBadRequestResponse("Invalid password. Should be base64 encoded")
	}

	if string(passwordString) != service.config.AdminPassword() {
		return responses.NewUnauthorizedResponse("Invalid password")
	}

	exp := time.Now().Add(service.config.JWTExpiration()).Unix()

	token := service.jwtService.GenerateToken(exp)

	service.sessionRepository.Create(ctx, models.Session{
		Token:     token,
		ClientIp:  clientIp,
		ExpiresAt: time.Unix(exp, 0),
	})

	return responses.NewSuccessResponse(token)
}

func (service *AuthService) Verify(ctx context.Context, clientIp, token string) *responses.Response {
	err := service.jwtService.VerifyToken(token)
	if err != nil {
		return responses.NewUnauthorizedResponse(err.Error())
	}

	session := service.sessionRepository.GetByToken(ctx, token)

	if session == nil {
		return responses.NewUnauthorizedResponse("Invalid token")
	}

	if session.ClientIp != clientIp {
		return responses.NewUnauthorizedResponse("Invalid token")
	}

	return responses.NewSuccessResponse(nil)
}

func (service *AuthService) Refresh(ctx context.Context, clientIp, token string) *responses.Response {
	err := service.jwtService.VerifyToken(token)
	if err != nil {
		return responses.NewUnauthorizedResponse("Invalid token: ")
	}

	session := service.sessionRepository.GetByToken(ctx, token)
	if session == nil {
		return responses.NewUnauthorizedResponse("Invalid session token")
	}

	if session.ClientIp != clientIp {
		return responses.NewUnauthorizedResponse("Invalid token")
	}

	newExp := time.Now().Add(service.config.JWTExpiration()).Unix()
	newToken := service.jwtService.GenerateToken(newExp)

	newSession := models.Session{
		Token:     newToken,
		ClientIp:  clientIp,
		ExpiresAt: time.Unix(newExp, 0),
	}
	service.sessionRepository.Create(ctx, newSession)
	service.sessionRepository.DeleteByToken(ctx, token)

	return responses.NewSuccessResponse(newToken)
}
