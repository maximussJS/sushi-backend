package models

import (
	"gorm.io/gorm"
	"sushi-backend/utils"
	"time"
)

type CategoryModel struct {
	Id          string         `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:255;not null;unique" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Products    []ProductModel `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	CreatedAt   time.Time      `json:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" json:"updatedAt"`
}

func (c *CategoryModel) TableName() string {
	return "categories"
}

func (c *CategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.Id = utils.NewUUID()
	return
}
