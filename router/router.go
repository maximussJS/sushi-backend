package router

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/constants"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/internal/logger"
	"sushi-backend/pkg/rate_limit"
)

type Router struct {
	router                 *mux.Router
	logger                 logger.ILogger
	config                 config.IConfig
	ipRateLimiter          rate_limit.IIpRateLimiter
	orderController        interfaces.IOrderController
	orderFlowController    interfaces.IOrderFlowController
	categoryController     interfaces.ICategoryController
	productController      interfaces.IProductController
	productImageController interfaces.IProductImageController
	analyticController     interfaces.IAnalyticController
}

func NewRouter(deps RouterDependencies) *Router {
	r := &Router{
		router:                 mux.NewRouter(),
		logger:                 deps.Logger,
		config:                 deps.Config,
		ipRateLimiter:          deps.IPRateLimiter,
		orderController:        deps.OrderController,
		orderFlowController:    deps.OrderFlowController,
		categoryController:     deps.CategoryController,
		productController:      deps.ProductController,
		productImageController: deps.ProductImageController,
		analyticController:     deps.AnalyticController,
	}

	orderRouter := r.router.PathPrefix("/api/v1/orders").Subrouter()

	orderRouter.HandleFunc("", r.noCache(r.wrapResponse(r.orderController.Create))).Methods("POST")
	orderRouter.HandleFunc("", r.wrapResponse(r.orderController.GetAll)).Methods("GET")
	orderRouter.HandleFunc("/{id}", r.wrapResponse(r.orderController.GetById)).Methods("GET")
	orderRouter.HandleFunc("/{id}", r.wrapResponse(r.orderController.DeleteById)).Methods("DELETE")

	r.addDefaultMiddlewares(orderRouter)

	orderFlowRouter := r.router.PathPrefix("/api/v1/order-flow").Subrouter()

	orderFlowRouter.HandleFunc("/{id}/{estimatedTimeInMs}/start-processing", r.noCache(r.wrapResponse(r.orderFlowController.StartProcessing))).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/ready-to-deliver", r.noCache(r.wrapResponse(r.orderFlowController.ReadyToDeliver))).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/{estimatedTimeInMs}/start-delivering", r.noCache(r.wrapResponse(r.orderFlowController.StartDelivering))).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/delivered", r.noCache(r.wrapResponse(r.orderFlowController.Delivered))).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/cancel", r.noCache(r.wrapResponse(r.orderFlowController.Cancel))).Methods("POST")

	r.addDefaultMiddlewares(orderFlowRouter)

	categoryRouter := r.router.PathPrefix("/api/v1/categories").Subrouter()

	categoryRouter.HandleFunc("", r.wrapResponse(r.categoryController.GetAll)).Methods("GET")
	categoryRouter.HandleFunc("", r.noCache(r.wrapResponse(r.categoryController.Create))).Methods("POST")
	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.DeleteById)).Methods("DELETE")
	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.GetById)).Methods("GET")
	categoryRouter.HandleFunc("/{id}", r.noCache(r.wrapResponse(r.categoryController.UpdateById))).Methods("PATCH")

	r.addDefaultMiddlewares(categoryRouter)

	productRouter := r.router.PathPrefix("/api/v1/products").Subrouter()

	productRouter.HandleFunc("", r.wrapResponse(r.productController.GetAll)).Methods("GET")
	productRouter.HandleFunc("", r.noCache(r.wrapResponse(r.productController.Create))).Methods("POST")
	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.DeleteById)).Methods("DELETE")
	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.GetById)).Methods("GET")
	productRouter.HandleFunc("/{id}", r.noCache(r.wrapResponse(r.productController.UpdateById))).Methods("PATCH")

	r.addDefaultMiddlewares(productRouter)

	productImageRouter := r.router.PathPrefix("/api/v1/product-images").Subrouter()

	productImageRouter.HandleFunc("/{id}", r.noCache(r.wrapResponse(r.productImageController.Create))).Methods("POST")
	productImageRouter.HandleFunc("/{id}", r.wrapResponse(r.productImageController.DeleteById)).Methods("DELETE")
	productImageRouter.HandleFunc("/{id}", r.noCache(r.wrapResponse(r.productImageController.GetById))).Methods("GET")

	r.addDefaultMiddlewares(productImageRouter)

	analyticRouter := r.router.PathPrefix("/api/v1/analytics").Subrouter()

	analyticRouter.HandleFunc("/orders/{startTimeInMs}", r.noCache(r.wrapResponse(r.analyticController.GetOrdersAnalytic))).Methods("GET")
	analyticRouter.HandleFunc("/products/{startTimeInMs}", r.noCache(r.wrapResponse(r.analyticController.GetTopOrderedProducts))).Methods("GET")

	r.addDefaultMiddlewares(analyticRouter)

	return r
}

func (router *Router) GetRouter() http.Handler {
	return router.withCors()
}

func (router *Router) withCors() http.Handler {
	switch router.config.AppEnv() {
	case constants.DevelopmentEnv:
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		})

		return c.Handler(router.router)
	case constants.ProductionEnv:
		c := cors.New(cors.Options{
			AllowedOrigins:   router.config.AllowedOrigins(),
			AllowedHeaders:   router.config.AllowedHeaders(),
			AllowedMethods:   router.config.AllowedMethods(),
			AllowCredentials: router.config.AllowCredentials(),
		})

		return c.Handler(router.router)
	default:
		panic("Unknown environment")
	}
}
