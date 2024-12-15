package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/pkg/logger"
	"sushi-backend/services/interfaces"
)

type ProductControllerDependencies struct {
	dig.In

	Logger         logger.ILogger             `name:"Logger"`
	ProductService interfaces.IProductService `name:"ProductService"`
}
