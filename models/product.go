package models

import (
	"gorm.io/gorm"
	"sushi-backend/utils"
	"time"
)

type ProductModel struct {
	Id          string              `gorm:"primaryKey" json:"id"`
	Name        string              `gorm:"size:255;not null;unique_index" json:"name"`
	Description string              `gorm:"type:text" json:"description,omitempty"`
	Price       float64             `gorm:"not null" json:"price"`
	CategoryId  string              `gorm:"not null;index" json:"categoryId"`
	Images      []ProductImageModel `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"images,omitempty"`
	Category    CategoryModel       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
}

func (p *ProductModel) TableName() string {
	return "products"
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = utils.NewUUID()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return
}
