package http_server

import (
	"context"
	"fmt"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
	"sushi-backend/router"
	"time"
)

type HttpServer struct {
	logger logger.ILogger
	config config.IConfig
	router *router.Router
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.logger.Log("Calling shutdown on http server")

	return nil
}

func StartHttpServer(deps HttpServerDependencies) {
	defer deps.ShutdownWaitGroup.Done()

	server := &HttpServer{
		logger: deps.Logger,
		config: deps.Config,
		router: deps.Router,
	}

	port := server.config.HttpPort()

	server.logger.Log(fmt.Sprintf("Starting http server on port %s", port))

	go func() {
		if err := http.ListenAndServe(port, server.router.GetRouter()); err != nil && err != http.ErrServerClosed {
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
