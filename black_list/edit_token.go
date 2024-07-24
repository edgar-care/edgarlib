package black_list

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("my_secret_key")

func modifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		claims["status"] = false
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		newTokenString, err := newToken.SignedString(secretKey)
		if err != nil {
			return "", err
		}
		return newTokenString, nil
	}

	return "", fmt.Errorf("could not parse claims")
}
