package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
)

type ProductImageControllerDependencies struct {
	dig.In

	Logger              logger.ILogger                  `name:"Logger"`
	Config              config.IConfig                  `name:"Config"`
	ProductImageService interfaces.IProductImageService `name:"ProductImageService"`
}
