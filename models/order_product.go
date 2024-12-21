package models

import (
	"time"

	"gorm.io/gorm"
	"sushi-backend/utils"
)

type OrderProductModel struct {
	Id        string       `gorm:"primaryKey" json:"id"`
	OrderId   string       `gorm:"not null;index" json:"orderId"`
	ProductId string       `gorm:"not null;index" json:"productId"`
	Quantity  int          `gorm:"not null" json:"quantity"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Order     OrderModel   `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Product   ProductModel `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (op *OrderProductModel) TableName() string {
	return "order_products"
}

func (op *OrderProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	op.Id = utils.NewUUID()
	op.CreatedAt = time.Now()
	op.UpdatedAt = time.Now()
	return
}

func (op *OrderProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	op.UpdatedAt = time.Now()
	return
}
