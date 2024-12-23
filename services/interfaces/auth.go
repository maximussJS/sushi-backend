package interfaces

import (
	"sushi-backend/types/responses"
)

type IAuthService interface {
	Authorize(clientIp, passwordInBase64String string) *responses.Response
	Verify(clientIp, token string) *responses.Response
}
