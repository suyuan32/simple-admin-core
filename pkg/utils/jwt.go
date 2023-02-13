package utils

import (
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func NewJwtToken(secretKey, uuid, roleString string, iat, seconds int64, roleIds []string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = uuid
	claims[roleString] = strings.Join(roleIds, ",")
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
