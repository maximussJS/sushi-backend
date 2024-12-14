package http_server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sushi-backend/internal/interfaces/handlers"
	"sushi-backend/pkg/config"
	"sushi-backend/pkg/logger"
	"sushi-backend/pkg/rate_limit"
	"time"
)

type HttpServer struct {
	logger        logger.ILogger
	config        config.IConfig
	ipRateLimiter rate_limit.IIpRateLimiter
	orderHandler  handlers.IOrderHandler
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.logger.Log("Calling shutdown on http server")

	return nil
}

func StartHttpServer(deps HttpServerDependencies) {
	defer deps.ShutdownWaitGroup.Done()

	server := &HttpServer{
		logger:        deps.Logger,
		config:        deps.Config,
		ipRateLimiter: deps.IPRateLimiter,
		orderHandler:  deps.OrderHandler,
	}

	port := server.config.GetHttpPort()

	server.logger.Log(fmt.Sprintf("Starting http server on port %s", port))

	go func() {
		if err := http.ListenAndServe(port, server.createRouter()); err != nil && err != http.ErrServerClosed {
			server.logger.Fatal(fmt.Sprintf("Failed to start http server: %s", err))
		}
	}()

	select {
	case <-deps.ShutdownContext.Done():
		server.logger.Log("Shutting down HTTP server gracefully...")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			server.logger.Error(fmt.Sprintf("Failed to shutdown HTTP server gracefully: %s", err))
		}
	}

	server.logger.Log("HTTP server stopped")
}

func (s *HttpServer) createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders", s.orderHandler.CreateOrder).Methods("POST")

	router.Use(s.IPMiddleware)
	router.Use(s.logMiddleware)
	router.Use(s.limitMiddleware)

	return router
}
