package repositories

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sushi-backend/models"
	"sushi-backend/repositories/dependencies"
	"sushi-backend/utils"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(deps dependencies.SessionRepositoryDependencies) *SessionRepository {
	if deps.Config.RunMigration() {
		utils.PanicIfError(deps.DB.AutoMigrate(&models.Session{}))
	}

	return &SessionRepository{
		db: deps.DB,
	}
}

func (r *SessionRepository) Create(ctx context.Context, session models.Session) string {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Create(&session).Error)

	return session.Token
}

func (r *SessionRepository) GetByToken(ctx context.Context, token string) *models.Session {
	var session models.Session
	err := r.db.WithContext(ctx).Clauses(clause.Returning{}).Where("token = ?", token).First(&session).Error

	return utils.HandleRecordNotFound[*models.Session](&session, err)
}

func (r *SessionRepository) DeleteByToken(ctx context.Context, token string) {
	utils.PanicIfErrorIsNotContextError(r.db.WithContext(ctx).Where("token = ?", token).Delete(&models.Session{}).Error)
}
