package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/jwt"
)

type AuthServiceDependencies struct {
	dig.In

	Config     config.IConfig  `name:"Config"`
	JwtService jwt.IJwtService `name:"JwtService"`
}
