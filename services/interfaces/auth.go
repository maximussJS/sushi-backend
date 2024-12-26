package interfaces

import (
	"context"
	"sushi-backend/types/responses"
)

type IAuthService interface {
	Authorize(ctx context.Context, clientIp, passwordInBase64String string) *responses.Response
	Verify(ctx context.Context, clientIp, token string) *responses.Response
	Refresh(ctx context.Context, clientIp, token string) *responses.Response
}
