package router

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/pkg/logger"
	"sushi-backend/pkg/rate_limit"
)

type RouterDependencies struct {
	dig.In

	Logger             logger.ILogger                 `name:"Logger"`
	Config             config.IConfig                 `name:"Config"`
	OrderController    interfaces.IOrderController    `name:"OrderController"`
	CategoryController interfaces.ICategoryController `name:"CategoryController"`
	ProductController  interfaces.IProductController  `name:"ProductController"`
	IPRateLimiter      rate_limit.IIpRateLimiter      `name:"IpRateLimiter"`
}