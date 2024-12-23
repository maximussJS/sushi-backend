package dependencies

import (
	"go.uber.org/dig"
	"sushi-backend/services/interfaces"
)

type AuthControllerDependencies struct {
	dig.In

	AuthService interfaces.IAuthService `name:"AuthService"`
}
