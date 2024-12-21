package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/internal/logger"
	"sushi-backend/pkg/rate_limit"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
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
	}

	orderRouter := r.router.PathPrefix("/api/v1/orders").Subrouter()

	orderRouter.HandleFunc("", r.wrapResponse(r.orderController.Create)).Methods("POST")
	orderRouter.HandleFunc("", r.wrapResponse(r.orderController.GetAll)).Methods("GET")
	orderRouter.HandleFunc("/{id}", r.wrapResponse(r.orderController.GetById)).Methods("GET")
	orderRouter.HandleFunc("/{id}", r.wrapResponse(r.orderController.DeleteById)).Methods("DELETE")

	r.addDefaultMiddlewares(orderRouter)

	orderFlowRouter := r.router.PathPrefix("/api/v1/order-flow").Subrouter()

	orderFlowRouter.HandleFunc("/{id}/start-processing", r.wrapResponse(r.orderFlowController.StartProcessing)).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/ready-to-deliver", r.wrapResponse(r.orderFlowController.ReadyToDeliver)).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/start-delivering", r.wrapResponse(r.orderFlowController.StartDelivering)).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/delivered", r.wrapResponse(r.orderFlowController.Delivered)).Methods("POST")
	orderFlowRouter.HandleFunc("/{id}/cancel", r.wrapResponse(r.orderFlowController.Cancel)).Methods("POST")

	r.addDefaultMiddlewares(orderFlowRouter)

	categoryRouter := r.router.PathPrefix("/api/v1/categories").Subrouter()

	categoryRouter.HandleFunc("", r.wrapResponse(r.categoryController.GetAll)).Methods("GET")
	categoryRouter.HandleFunc("", r.wrapResponse(r.categoryController.Create)).Methods("POST")
	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.DeleteById)).Methods("DELETE")
	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.GetById)).Methods("GET")
	categoryRouter.HandleFunc("/{id}", r.wrapResponse(r.categoryController.UpdateById)).Methods("PATCH")

	r.addDefaultMiddlewares(categoryRouter)

	productRouter := r.router.PathPrefix("/api/v1/products").Subrouter()

	productRouter.HandleFunc("", r.wrapResponse(r.productController.GetAll)).Methods("GET")
	productRouter.HandleFunc("", r.wrapResponse(r.productController.Create)).Methods("POST")
	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.DeleteById)).Methods("DELETE")
	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.GetById)).Methods("GET")
	productRouter.HandleFunc("/{id}", r.wrapResponse(r.productController.UpdateById)).Methods("PATCH")

	r.addDefaultMiddlewares(productRouter)

	productImageRouter := r.router.PathPrefix("/api/v1/product-images").Subrouter()

	productImageRouter.HandleFunc("/{id}", r.wrapResponse(r.productImageController.Create)).Methods("POST")
	productImageRouter.HandleFunc("/{id}", r.wrapResponse(r.productImageController.DeleteById)).Methods("DELETE")
	productImageRouter.HandleFunc("/{id}", r.wrapResponse(r.productImageController.GetById)).Methods("GET")

	r.addDefaultMiddlewares(productImageRouter)

	return r
}

func (router *Router) GetRouter() *mux.Router {
	return router.router
}

type wrappedFn func(w http.ResponseWriter, r *http.Request) *responses.Response

func (router *Router) wrapResponse(fn wrappedFn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := fn(w, r)

		if resp.IsError() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			utils.PanicIfErrorWithResult(w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, resp.Msg))))
			return
		}

		if resp.Data != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)
			utils.PanicIfError(json.NewEncoder(w).Encode(&resp.Data))
			return
		}

		w.WriteHeader(resp.Status)
	}
}
