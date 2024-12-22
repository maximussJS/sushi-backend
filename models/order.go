package models

import (
	"sushi-backend/constants"
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	Id              uint                  `gorm:"unique;primaryKey;autoIncrement" json:"id" `
	FirstName       string                `gorm:"size:100;not null" json:"firstName"`
	LastName        string                `gorm:"size:100;not null" json:"lastName"`
	Price           float64               `gorm:"not null" json:"price"`
	Phone           string                `gorm:"size:20;not null" json:"phone"`
	DeliveryAddress string                `gorm:"size:255;not null" json:"deliveryAddress"`
	PaymentMethod   string                `gorm:"size:50;not null" json:"paymentMethod"`
	OrderedProducts []OrderProductModel   `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"orderedProducts"`
	Status          constants.OrderStatus `gorm:"type:varchar(50);not null;" json:"status"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
	DeliveredAt     time.Time             `gorm:"type:TIMESTAMP;null;default:null" json:"deliveredAt"`
	ProcessedAt     time.Time             `gorm:"type:TIMESTAMP;null;default:null" json:"processedAt"`
}

func (o *OrderModel) TableName() string {
	return "orders"
}

func (o *OrderModel) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
	o.Status = constants.StatusCreated
	return
}

func (o *OrderModel) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedAt = time.Now()
	return
}
