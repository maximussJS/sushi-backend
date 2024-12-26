package services

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sushi-backend/config"
	"sushi-backend/internal/telegram"
	"sushi-backend/models"
	"sushi-backend/repositories/interfaces"
	dependencies "sushi-backend/services/dependecies"
	"sushi-backend/types/requests"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
)

type OrderService struct {
	config            config.IConfig
	telegram          telegram.ITelegram
	orderRepository   interfaces.IOrderRepository
	productRepository interfaces.IProductRepository
}

func NewOrderService(deps dependencies.OrderServiceDependencies) *OrderService {
	return &OrderService{
		config:            deps.Config,
		telegram:          deps.Telegram,
		orderRepository:   deps.OrderRepository,
		productRepository: deps.ProductRepository,
	}
}

func (o *OrderService) Create(request requests.CreateOrderRequest) *responses.Response {
	productIds := make([]string, len(request.OrderedProducts))

	for i, orderedProduct := range request.OrderedProducts {
		productIds[i] = orderedProduct.ProductId
	}

	productsMap, err := o.getProductsMap(productIds)
	if err != nil {
		return responses.NewBadRequestResponse(err.Error())
	}

	orderId := o.orderRepository.Create(request.ToOrderModel(productsMap))

	newOrder := o.orderRepository.GetById(orderId)

	msg := o.constructNewOrderTelegramMessage(newOrder)

	o.telegram.SendMessageToChannel(o.config.TelegramOrdersChatId(), msg, true)

	return responses.NewSuccessResponse(newOrder)
}

func (o *OrderService) GetById(id uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	return responses.NewSuccessResponse(order)
}

func (o *OrderService) GetAll(limit, offset int) *responses.Response {
	orders := o.orderRepository.GetAll(limit, offset)

	return responses.NewSuccessResponse(orders)
}

func (o *OrderService) DeleteById(id uint) *responses.Response {
	order := o.orderRepository.GetById(id)

	if order == nil {
		return responses.NewNotFoundResponse(fmt.Sprintf("Order with id %d not found", id))
	}

	o.orderRepository.DeleteById(id)

	return responses.NewSuccessResponse(order)
}

func (o *OrderService) getProductsMap(productsIds []string) (map[string]models.ProductModel, error) {
	products := o.productRepository.GetByIds(productsIds)

	productsMap := make(map[string]models.ProductModel, len(products))
	for _, product := range products {
		productsMap[product.Id] = product
	}

	for _, productId := range productsIds {
		if _, ok := productsMap[productId]; !ok {
			return nil, errors.New(fmt.Sprintf("Product with id %s not found", productId))
		}
	}

	return productsMap, nil
}

func (o *OrderService) constructNewOrderTelegramMessage(order *models.OrderModel) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("*New Order: %s*\n\n", utils.EscapeMarkdown(strconv.Itoa(int(order.Id)))))

	sb.WriteString("*Ordered Products:*\n")
	for _, op := range order.OrderedProducts {
		// List each product with quantity and name
		sb.WriteString(fmt.Sprintf("  - *%d* × %s\n", op.Quantity, utils.EscapeMarkdown(op.Product.Name)))
	}
	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf("*Price:* $%.2f\n", order.Price))

	sb.WriteString(fmt.Sprintf("*Delivery Address:* %s\n", utils.EscapeMarkdown(order.DeliveryAddress)))

	fullName := fmt.Sprintf("%s %s", utils.EscapeMarkdown(order.FirstName), utils.EscapeMarkdown(order.LastName))
	sb.WriteString(fmt.Sprintf("*Fullname:* %s\n", fullName))

	sb.WriteString(fmt.Sprintf("*Phone Number:* %s\n", utils.EscapeMarkdown(order.Phone)))

	sb.WriteString(fmt.Sprintf("*Payment Method:* %s\n", utils.EscapeMarkdown(order.PaymentMethod)))

	createdAt := order.CreatedAt.Format("2006-01-02 15:04:05")
	sb.WriteString(fmt.Sprintf("*Created At:* %s\n", createdAt))

	return sb.String()
}
