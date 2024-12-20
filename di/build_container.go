package di

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/controllers"
	controllers_interfaces "sushi-backend/controllers/interfaces"
	"sushi-backend/internal/cloudinary"
	"sushi-backend/internal/db"
	logger2 "sushi-backend/internal/logger"
	"sushi-backend/internal/telegram"
	"sushi-backend/internal/tmp_file_storage"
	"sushi-backend/pkg/rate_limit"
	"sushi-backend/repositories"
	repositories_interfaces "sushi-backend/repositories/interfaces"
	"sushi-backend/router"
	"sushi-backend/services"
	services_interfaces "sushi-backend/services/interfaces"
	"sushi-backend/utils"
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
		utils.PanicIfError(container.Provide(dependency.Constructor, dig.Name(dependency.Token)))
		return
	}

	utils.PanicIfError(container.Provide(
		dependency.Constructor,
		dig.As(dependency.Interface),
		dig.Name(dependency.Token),
	))
}

// GetInternalDependencies The list of internal dependencies that are required for the application to run.
func getInternalDependencies() []Dependency {
	return []Dependency{
		{
			Constructor: repositories.NewProductImageRepository,
			Interface:   new(repositories_interfaces.IProductImageRepository),
			Token:       "ProductImageRepository",
		},
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
			Constructor: services.NewProductImageService,
			Interface:   new(services_interfaces.IProductImageService),
			Token:       "ProductImageService",
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
			Constructor: services.NewOrderService,
			Interface:   new(services_interfaces.IOrderService),
			Token:       "OrderService",
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
			Constructor: controllers.NewProductImageController,
			Interface:   new(controllers_interfaces.IProductImageController),
			Token:       "ProductImageController",
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
			Constructor: logger2.NewLogger,
			Interface:   new(logger2.ILogger),
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
			Constructor: tmp_file_storage.NewTmpFileStorage,
			Interface:   new(tmp_file_storage.ITmpFileStorage),
			Token:       "TmpFileStorage",
		},
		{
			Constructor: cloudinary.NewCloudinary,
			Interface:   new(cloudinary.ICloudinary),
			Token:       "Cloudinary",
		},
		{
			Constructor: telegram.NewTelegram,
			Interface:   new(telegram.ITelegram),
			Token:       "Telegram",
		},
		{
			Constructor: db.NewDB,
			Interface:   nil,
			Token:       "DB",
		},
	}
}
