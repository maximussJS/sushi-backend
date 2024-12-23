package jwt

type IJwtService interface {
	GenerateTokenWithClientIp(clientIp string) string
	VerifyTokenWithClientIp(tokenString string) (clientIP string, err error)
}
