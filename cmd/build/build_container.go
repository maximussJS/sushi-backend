package build

import (
	"go.uber.org/dig"
	"sushi-backend/common/types"
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

func AppendDependenciesToContainer(container *dig.Container, dependencies []types.Dependency) *dig.Container {
	for _, dep := range dependencies {
		mustProvideDependency(container, dep)
	}

	return container
}

func mustProvideDependency(container *dig.Container, dependency types.Dependency) {
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
func getInternalDependencies() []types.Dependency {
	return []types.Dependency{
		{
			Constructor: handlers.NewOrderHandler,
			Interface:   new(handlers_interfaces.IOrderHandler),
			Token:       "OrderHandler",
		},
	}
}

// getRequiredDependencies The list of dependencies that are required for the application to run.
func getRequiredDependencies() []types.Dependency {
	return []types.Dependency{
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
