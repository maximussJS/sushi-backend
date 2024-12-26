package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	Token     string    `gorm:"primaryKey" json:"id"`
	ClientIp  string    `gorm:"size:255;not null" json:"clientIp"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `gorm:"type:TIMESTAMP;null;default:null" json:"expiresAt"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return
}
