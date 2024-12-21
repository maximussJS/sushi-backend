package requests

import (
	"fmt"
	"sushi-backend/models"
)

type OrderedProduct struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,min=1"`
}

type CreateOrderRequest struct {
	FirstName       string           `json:"firstName" validate:"required,min=1,max=100"`
	LastName        string           `json:"lastName" validate:"required,min=1,max=100"`
	Phone           string           `json:"phone" validate:"required,min=5,max=20"`
	DeliveryAddress string           `json:"deliveryAddress" validate:"required,min=5,max=255"`
	PaymentMethod   string           `json:"paymentMethod" validate:"required,min=3,max=50"`
	OrderedProducts []OrderedProduct `json:"orderedProducts" validate:"required,min=1,max=100,dive,required"`
}

func (r CreateOrderRequest) ToOrderModel(productsMap map[string]models.ProductModel) models.OrderModel {
	orderedProducts := make([]models.OrderProductModel, len(r.OrderedProducts))
	price := 0.0
	for i, orderedProduct := range r.OrderedProducts {
		product, ok := productsMap[orderedProduct.ProductId]
		if !ok {
			panic(fmt.Sprintf("Product with id %s not found. Error in validation", orderedProduct.ProductId))
		}
		orderedProducts[i] = models.OrderProductModel{
			ProductId: product.Id,
			Quantity:  orderedProduct.Quantity,
		}

		price += product.Price * float64(orderedProduct.Quantity)
	}
	return models.OrderModel{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		Price:           price,
		Phone:           r.Phone,
		DeliveryAddress: r.DeliveryAddress,
		PaymentMethod:   r.PaymentMethod,
		OrderedProducts: orderedProducts,
	}
}
