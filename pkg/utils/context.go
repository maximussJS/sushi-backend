package utils

import (
	"context"
	"sushi-backend/pkg/constants"
)

func GetContextWithClientIp(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, constants.ClientIpKey, ip)
}

func GetClientIpFromContext(ctx context.Context) string {
	return ctx.Value(constants.ClientIpKey).(string)
}
