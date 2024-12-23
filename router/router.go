package router

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/constants"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/internal/jwt"
	"sushi-backend/internal/logger"
	"sushi-backend/internal/rate_limit"
)

type Router struct {
	router                 *mux.Router
	logger                 logger.ILogger
	config                 config.IConfig
	ipRateLimiter          rate_limit.IIpRateLimiter
	jwtService             jwt.IJwtService
	orderController        interfaces.IOrderController
	orderFlowController    interfaces.IOrderFlowController
	categoryController     interfaces.ICategoryController
	productController      interfaces.IProductController
	productImageController interfaces.IProductImageController
	analyticController     interfaces.IAnalyticController
	authController         interfaces.IAuthController
}

func NewRouter(deps RouterDependencies) *Router {
	r := &Router{
		router:                 mux.NewRouter(),
		logger:                 deps.Logger,
		config:                 deps.Config,
		ipRateLimiter:          deps.IPRateLimiter,
		jwtService:             deps.JwtService,
		orderController:        deps.OrderController,
		orderFlowController:    deps.OrderFlowController,
		categoryController:     deps.CategoryController,
		productController:      deps.ProductController,
		productImageController: deps.ProductImageController,
		analyticController:     deps.AnalyticController,
		authController:         deps.AuthController,
	}

	authRouter := r.router.PathPrefix("/api/v1/auth").Subrouter()

	r.addDefaultMiddlewares(authRouter)

	authRouter.HandleFunc("", r.noCache(r.wrapResponse(r.authController.Authorize))).Methods("POST")
	authRouter.HandleFunc("/verify", r.noCache(r.wrapResponse(r.authController.Verify))).Methods("POST")

	orderRouter := r.router.PathPrefix("/api/v1/orders").Subrouter()

	r.addDefaultMiddlewares(orderRouter)

	orderRouter.HandleFunc("",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderController.Create)))).Methods("POST")

	orderRouter.HandleFunc("", r.wrapResponse(r.orderController.GetAll)).Methods("GET")

	orderRouter.HandleFunc("/{id}", r.wrapResponse(r.orderController.GetById)).Methods("GET")

	orderRouter.HandleFunc("/{id}", r.noCache(r.isAdmin(r.wrapResponse(r.orderController.DeleteById)))).Methods("DELETE")

	orderFlowRouter := r.router.PathPrefix("/api/v1/order-flow").Subrouter()

	r.addDefaultMiddlewares(orderFlowRouter)

	orderFlowRouter.HandleFunc("/{id}/{estimatedTimeInMs}/start-processing",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderFlowController.StartProcessing)))).Methods("POST")

	orderFlowRouter.HandleFunc("/{id}/ready-to-deliver",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderFlowController.ReadyToDeliver)))).Methods("POST")

	orderFlowRouter.HandleFunc("/{id}/{estimatedTimeInMs}/start-delivering",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderFlowController.StartDelivering)))).Methods("POST")

	orderFlowRouter.HandleFunc("/{id}/delivered",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderFlowController.Delivered)))).Methods("POST")

	orderFlowRouter.HandleFunc("/{id}/cancel",
		r.noCache(r.isAdmin(r.wrapResponse(r.orderFlowController.Cancel)))).Methods("POST")

	categoryRouter := r.router.PathPrefix("/api/v1/categories").Subrouter()

	r.addDefaultMiddlewares(categoryRouter)

	categoryRouter.HandleFunc("", r.wrapResponse(r.categoryController.GetAll)).Methods("GET")

	categoryRouter.HandleFunc("",
		r.noCache(r.isAdmin(r.wrapResponse(r.categoryController.Create)))).Methods("POST")

	categoryRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.categoryController.DeleteById)))).Methods("DELETE")

	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.GetById)).Methods("GET")

	categoryRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.categoryController.UpdateById)))).Methods("PATCH")

	productRouter := r.router.PathPrefix("/api/v1/products").Subrouter()

	r.addDefaultMiddlewares(productRouter)

	productRouter.HandleFunc("", r.wrapResponse(r.productController.GetAll)).Methods("GET")

	productRouter.HandleFunc("", r.noCache(r.isAdmin(r.wrapResponse(r.productController.Create)))).Methods("POST")

	productRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.productController.DeleteById)))).Methods("DELETE")

	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.GetById)).Methods("GET")

	productRouter.HandleFunc("/{id}", r.noCache(r.isAdmin(r.wrapResponse(r.productController.UpdateById)))).Methods("PATCH")

	productImageRouter := r.router.PathPrefix("/api/v1/product-images").Subrouter()

	r.addDefaultMiddlewares(productImageRouter)

	productImageRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.productImageController.Create)))).Methods("POST")

	productImageRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.productImageController.DeleteById)))).Methods("DELETE")

	productImageRouter.HandleFunc("/{id}",
		r.noCache(r.isAdmin(r.wrapResponse(r.productImageController.GetById)))).Methods("GET")

	analyticRouter := r.router.PathPrefix("/api/v1/analytics").Subrouter()

	r.addDefaultMiddlewares(analyticRouter)

	analyticRouter.HandleFunc("/orders/{startTimeInMs}", r.noCache(r.isAdmin(r.wrapResponse(r.analyticController.GetOrdersAnalytic)))).Methods("GET")
	analyticRouter.HandleFunc("/products/{startTimeInMs}", r.noCache(r.isAdmin(r.wrapResponse(r.analyticController.GetTopOrderedProducts)))).Methods("GET")

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
