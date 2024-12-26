package interfaces

import (
	"context"
	"sushi-backend/models"
)

type ISessionRepository interface {
	Create(ctx context.Context, session models.Session) string
	GetByToken(ctx context.Context, token string) *models.Session
	DeleteByToken(ctx context.Context, token string)
}
