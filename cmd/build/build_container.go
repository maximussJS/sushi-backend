package build

import (
	"go.uber.org/dig"
	"sushi-backend/internal/handlers"
	handlers_interfaces "sushi-backend/internal/interfaces/handlers"
	"sushi-backend/pkg/config"
	"sushi-backend/pkg/logger"
	"sushi-backend/pkg/rate_limit"
)

func BuildContainer() *dig.Container {
	c := dig.New()

	c = AppendDependenciesToContainer(c, getRequiredDependencies())
	c = AppendDependenciesToContainer(c, getInternalDependencies())

	return c
}

func AppendDependenciesToContainer(container *dig.Container, dependencies []Dependency) *dig.Container {
	for _, dep := range dependencies {
		mustProvideDependency(container, dep)
	}

	return container
}

func mustProvideDependency(container *dig.Container, dependency Dependency) {
	if dependency.Interface == nil {
		err := container.Provide(dependency.Constructor, dig.Name(dependency.Token))

		if err != nil {
			panic(err)
		}

		return
	}

	err := container.Provide(
		dependency.Constructor,
		dig.As(dependency.Interface),
		dig.Name(dependency.Token),
	)

	if err != nil {
		panic(err)
	}
}

// GetInternalDependencies The list of internal dependencies that are required for the application to run.
func getInternalDependencies() []Dependency {
	return []Dependency{
		{
			Constructor: handlers.NewOrderHandler,
			Interface:   new(handlers_interfaces.IOrderHandler),
			Token:       "OrderHandler",
		},
	}
}

// getRequiredDependencies The list of dependencies that are required for the application to run.
func getRequiredDependencies() []Dependency {
	return []Dependency{
		{
			Constructor: logger.NewLogger,
			Interface:   new(logger.ILogger),
			Token:       "Logger",
		},
		{
			Constructor: config.NewConfig,
			Interface:   new(config.IConfig),
			Token:       "Config",
		},
		{
			Constructor: rate_limit.NewIPRateLimiter,
			Interface:   new(rate_limit.IIpRateLimiter),
			Token:       "IpRateLimiter",
		},
	}
}
