package router

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/internal/logger"
	"sushi-backend/internal/rate_limit"
	services_interfaces "sushi-backend/services/interfaces"
)

type RouterDependencies struct {
	dig.In

	Logger                 logger.ILogger                     `name:"Logger"`
	Config                 config.IConfig                     `name:"Config"`
	OrderController        interfaces.IOrderController        `name:"OrderController"`
	OrderFlowController    interfaces.IOrderFlowController    `name:"OrderFlowController"`
	CategoryController     interfaces.ICategoryController     `name:"CategoryController"`
	ProductController      interfaces.IProductController      `name:"ProductController"`
	ProductImageController interfaces.IProductImageController `name:"ProductImageController"`
	AnalyticController     interfaces.IAnalyticController     `name:"AnalyticController"`
	AuthController         interfaces.IAuthController         `name:"AuthController"`
	IPRateLimiter          rate_limit.IIpRateLimiter          `name:"IpRateLimiter"`
	AuthService            services_interfaces.IAuthService   `name:"AuthService"`
}
