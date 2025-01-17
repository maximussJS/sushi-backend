package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"sushi-backend/config"
	"sushi-backend/utils"
	"time"
)

type JwtService struct {
	config config.IConfig
}

func NewJwtService(deps JwtServiceDependices) *JwtService {
	return &JwtService{config: deps.Config}
}

func (j *JwtService) GenerateToken(exp int64) string {
	claims := jwt.MapClaims{
		"exp": exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString := utils.PanicIfErrorWithResultReturning(token.SignedString(j.config.JWTSecretKey()))

	return tokenString
}

func (j *JwtService) VerifyToken(tokenString string) (err error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return j.config.JWTSecretKey(), nil
	})

	if err != nil {
		return errors.New("invalid token")
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	expiredAt := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().Unix() > expiredAt.Unix() {
		return errors.New("token expired")
	}

	return nil
}
