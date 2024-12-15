package models

import (
	"time"
)

type ProductModel struct {
	Id          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"size:255;not null;unique_index" json:"name"`
	Description string        `gorm:"type:text" json:"description,omitempty"`
	Price       float64       `gorm:"not null" json:"price"`
	CategoryID  uint          `gorm:"not null;index" json:"category_id"`
	Category    CategoryModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

func (p *ProductModel) TableName() string {
	return "products"
}
