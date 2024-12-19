package models

import (
	"gorm.io/gorm"
	"sushi-backend/utils"
	"time"
)

type ProductImageModel struct {
	Id                 string       `gorm:"primaryKey" json:"id"`
	ProductId          string       `gorm:"not null;index" json:"productId"`
	CloudinaryPublicId string       `gorm:"size:255;not null" json:"-"`
	Url                string       `gorm:"size:1024;not null" json:"url"`
	CreatedAt          time.Time    `json:"createdAt"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	Product            ProductModel `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (pi *ProductImageModel) TableName() string {
	return "product_images"
}

func (pi *ProductImageModel) BeforeCreate(tx *gorm.DB) (err error) {
	pi.Id = utils.NewUUID()
	pi.CreatedAt = time.Now()
	pi.UpdatedAt = time.Now()
	return
}

func (pi *ProductImageModel) BeforeUpdate(tx *gorm.DB) (err error) {
	pi.UpdatedAt = time.Now()
	return
}
