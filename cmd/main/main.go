package main

import (
	"context"
	"go.uber.org/dig"
	"os"
	"os/signal"
	"sushi-backend/cmd/build"
	"sushi-backend/common/types"
	"sushi-backend/pkg/http_server"
	"sync"
	"syscall"
)

func runServer(container *dig.Container) {
	if err := container.Invoke(http_server.StartHttpServer); err != nil {
		panic(err)
	}
}

func main() {
	shutdownContext, cancel := context.WithCancel(context.Background())

	defer cancel()

	container := build.BuildContainer()

	var wg sync.WaitGroup

	container = build.AppendDependenciesToContainer(container, []types.Dependency{
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
