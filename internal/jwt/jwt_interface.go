package jwt

type IJwtService interface {
	GenerateToken(exp int64) string
	VerifyToken(tokenString string) (err error)
}
