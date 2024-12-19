package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/internal/cloudinary"
	"sushi-backend/repositories/interfaces"
)

type ProductImageServiceDependencies struct {
	dig.In

	Cloudinary             cloudinary.ICloudinary             `name:"Cloudinary"`
	ProductRepository      interfaces.IProductRepository      `name:"ProductRepository"`
	ProductImageRepository interfaces.IProductImageRepository `name:"ProductImageRepository"`
}
