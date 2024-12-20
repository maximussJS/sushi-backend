package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/logger"
	"sushi-backend/services/interfaces"
)

type ProductImageControllerDependencies struct {
	dig.In

	Logger              logger.ILogger                  `name:"Logger"`
	ProductImageService interfaces.IProductImageService `name:"ProductImageService"`
}
