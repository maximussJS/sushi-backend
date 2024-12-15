package di

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/controllers"
	controllers_interfaces "sushi-backend/controllers/interfaces"
	"sushi-backend/internal/db"
	"sushi-backend/pkg/logger"
	"sushi-backend/pkg/rate_limit"
	"sushi-backend/repositories"
	repositories_interfaces "sushi-backend/repositories/interfaces"
	"sushi-backend/router"
	"sushi-backend/services"
	services_interfaces "sushi-backend/services/interfaces"
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
			Constructor: repositories.NewProductRepository,
			Interface:   new(repositories_interfaces.IProductRepository),
			Token:       "ProductRepository",
		},
		{
			Constructor: repositories.NewCategoryRepository,
			Interface:   new(repositories_interfaces.ICategoryRepository),
			Token:       "CategoryRepository",
		},
		{
			Constructor: services.NewProductService,
			Interface:   new(services_interfaces.IProductService),
			Token:       "ProductService",
		},
		{
			Constructor: services.NewCategoryService,
			Interface:   new(services_interfaces.ICategoryService),
			Token:       "CategoryService",
		},
		{
			Constructor: controllers.NewOrderController,
			Interface:   new(controllers_interfaces.IOrderController),
			Token:       "OrderController",
		},
		{
			Constructor: controllers.NewCategoryController,
			Interface:   new(controllers_interfaces.ICategoryController),
			Token:       "CategoryController",
		},
		{
			Constructor: controllers.NewProductController,
			Interface:   new(controllers_interfaces.IProductController),
			Token:       "ProductController",
		},
		{
			Constructor: router.NewRouter,
			Interface:   nil,
			Token:       "Router",
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
		{
			Constructor: db.NewDB,
			Interface:   nil,
			Token:       "DB",
		},
	}
}
