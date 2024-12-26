package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/config"
	"sushi-backend/internal/jwt"
	"sushi-backend/repositories/interfaces"
)

type AuthServiceDependencies struct {
	dig.In

	Config            config.IConfig                `name:"Config"`
	JwtService        jwt.IJwtService               `name:"JwtService"`
	SessionRepository interfaces.ISessionRepository `name:"SessionRepository"`
}
