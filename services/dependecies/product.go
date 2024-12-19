package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/repositories/interfaces"
)

type ProductServiceDependencies struct {
	dig.In

	ProductRepository  interfaces.IProductRepository  `name:"ProductRepository"`
	CategoryRepository interfaces.ICategoryRepository `name:"CategoryRepository"`
}
