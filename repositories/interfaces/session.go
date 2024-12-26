package interfaces

import "sushi-backend/models"

type ISessionRepository interface {
	Create(session models.Session) string
	GetByToken(token string) *models.Session
	DeleteByToken(token string)
}
