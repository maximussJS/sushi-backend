package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/repositories/interfaces"
)

type CategoryServiceDependencies struct {
	dig.In

	CategoryRepository interfaces.ICategoryRepository `name:"CategoryRepository"`
}
