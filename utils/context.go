package utils

import (
	"context"
	"sushi-backend/constants"
)

func GetContextWithClientIp(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, constants.ClientIpKey, ip)
}

func GetClientIpFromContext(ctx context.Context) string {
	return ctx.Value(constants.ClientIpKey).(string)
}
