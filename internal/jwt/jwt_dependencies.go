package jwt

import (
	"go.uber.org/dig"
	"sushi-backend/config"
)

type JwtServiceDependices struct {
	dig.In

	Config config.IConfig `name:"Config"`
}
