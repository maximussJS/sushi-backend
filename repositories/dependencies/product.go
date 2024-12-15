package dependencies

import (
	"go.uber.org/dig"
	"gorm.io/gorm"
	"sushi-backend/config"
)

type ProductRepositoryDependencies struct {
	dig.In

	DB     *gorm.DB       `name:"DB"`
	Config config.IConfig `name:"Config"`
}
