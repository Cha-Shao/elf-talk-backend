package utils

import (
	"elf-talk/config"

	"github.com/dgrijalva/jwt-go"
)

var Secret = []byte(config.Conf.Secret)

type Claims struct {
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

// 生成token
func Generate(payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}
	for key, value := range payload {
		claims[key] = value
	}
	claims["iss"] = "elf-talk"
	claims["exp"] = 0

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(Secret)

	if err != nil {
		return "", err
	}

	return token, nil
}

// 解析token
func Parse(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
