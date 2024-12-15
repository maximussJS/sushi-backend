package main

import (
	"context"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"sushi-backend/di"
	"sushi-backend/internal/http_server"
	"sushi-backend/utils"
	"sync"
	"syscall"
)

func runServer(container *dig.Container) {
	utils.PanicIfError(container.Invoke(http_server.StartHttpServer))
}

func main() {
	shutdownContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	container := di.BuildContainer()

	var wg sync.WaitGroup

	container = di.AppendDependenciesToContainer(container, []di.Dependency{
		{
			Constructor: func() context.Context {
				return shutdownContext
			},
			Interface: nil,
			Token:     "ShutdownContext",
		},
		{
			Constructor: func() *sync.WaitGroup {
				return &wg
			},
			Interface: nil,
			Token:     "ShutdownWaitGroup",
		},
	})

	wg.Add(1)
	go runServer(container)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	<-signalCh

	cancel()

	wg.Wait()
}
