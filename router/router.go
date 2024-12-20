package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/controllers/interfaces"
	"sushi-backend/internal/logger"
	"sushi-backend/pkg/rate_limit"
	"sushi-backend/types/responses"
	"sushi-backend/utils"
	"time"
)

type Router struct {
	router                 *mux.Router
	logger                 logger.ILogger
	config                 config.IConfig
	ipRateLimiter          rate_limit.IIpRateLimiter
	orderController        interfaces.IOrderController
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
		categoryController:     deps.CategoryController,
		productController:      deps.ProductController,
		productImageController: deps.ProductImageController,
	}

	orderRouter := r.router.PathPrefix("/api/v1/orders").Subrouter()

	orderRouter.HandleFunc("", r.wrapResponse(r.orderController.CreateOrder)).Methods("POST")

	r.addDefaultMiddlewares(orderRouter)

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

func (router *Router) limitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := utils.GetClientIpFromContext(r.Context())

		limiter := router.ipRateLimiter.GetLimiter(clientIP)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (router *Router) logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		clientIP := utils.GetClientIpFromContext(r.Context())

		router.logger.Log(fmt.Sprintf(
			"[%s] %s %s %s %.0fm%.0fs%dms%dns %s",
			start.Format(time.RFC3339),
			r.Method,
			r.RequestURI,
			r.Proto,
			duration.Minutes(),
			duration.Seconds(),
			duration.Milliseconds(),
			duration.Microseconds(),
			clientIP,
		))
	})
}

func (router *Router) iPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr
		if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
			clientIP = ip
		}

		next.ServeHTTP(w, r.WithContext(utils.GetContextWithClientIp(r.Context(), clientIP)))
	})
}

func (router *Router) GetRouter() *mux.Router {
	return router.router
}

func (router *Router) addDefaultMiddlewares(r *mux.Router) {
	r.Use(router.iPMiddleware)
	r.Use(router.limitMiddleware)
	r.Use(router.logMiddleware)
}

type wrappedFn func(w http.ResponseWriter, r *http.Request) *responses.Response

func (router *Router) wrapResponse(fn wrappedFn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := fn(w, r)

		if resp == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			utils.PanicIfErrorWithResult(w.Write([]byte(`{"message": "Internal server error"}`)))
			return
		}

		if resp.IsError() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.Status)

			if resp.Status == http.StatusInternalServerError {
				router.logger.Error(resp.Error.Error())
			}

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
