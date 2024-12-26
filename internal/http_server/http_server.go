package http_server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"sushi-backend/config"
	"sushi-backend/constants"
	"sushi-backend/internal/logger"
	"sushi-backend/router"
	"time"
)

type HttpServer struct {
	logger  logger.ILogger
	config  config.IConfig
	router  *router.Router
	_server http.Server
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.logger.Log("Calling shutdown on http server")

	return nil
}

func StartHttpServer(deps HttpServerDependencies) {
	defer deps.ShutdownWaitGroup.Done()

	tlsConfig := &tls.Config{
		ClientAuth: tls.NoClientCert,
		MinVersion: tls.VersionTLS11,
	}

	port := deps.Config.HttpPort()

	server := &HttpServer{
		logger: deps.Logger,
		config: deps.Config,
		router: deps.Router,
		_server: http.Server{
			Addr:         port,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  30 * time.Second,
			TLSConfig:    tlsConfig,
			Handler:      deps.Router.GetRouter(),
		},
	}

	go func() {
		if server.config.AppEnv() == constants.DevelopmentEnv {
			server.logger.Log(fmt.Sprintf("Starting http server on port %s", port))

			if err := server._server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				server.logger.Fatal(fmt.Sprintf("Failed to start http server: %s", err))
			}
		} else {
			server.logger.Log(fmt.Sprintf("Starting https server on port %s", port))
			if err := server._server.ListenAndServeTLS(server.config.SSLCertPath(), server.config.SSLKeyPath()); err != nil && err != http.ErrServerClosed {
				server.logger.Fatal(fmt.Sprintf("Failed to start https server: %s", err))
			}
		}
	}()

	select {
	case <-deps.ShutdownContext.Done():
		server.logger.Log("Shutting down server gracefully...")
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			server.logger.Error(fmt.Sprintf("Failed to shutdown HTTP server gracefully: %s", err))
		}
	}

	server.logger.Log("Server stopped")
}
