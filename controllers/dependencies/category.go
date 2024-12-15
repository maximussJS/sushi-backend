package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/pkg/logger"
	"sushi-backend/services/interfaces"
)

type CategoryControllerDependencies struct {
	dig.In

	Logger          logger.ILogger              `name:"Logger"`
	CategoryService interfaces.ICategoryService `name:"CategoryService"`
}
