package tools

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	CloudJWTKey string
)

func VerifyJWT(jwtToken string) (string, string, string, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(CloudJWTKey))
	if err != nil {
		return "", "", "", err
	}
	var publicKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return publicKey, nil }

	token, err := jwt.Parse(jwtToken, publicKeyFunc)
	if err != nil {
		return "", "", "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	exp := claims["exp"].(float64)
	userID := claims["sub"].(string)
	login := claims["login"].(string)
	name := claims["name"].(string)

	expiresAt := time.Unix(int64(exp), 0)

	if int64(exp) < time.Now().Unix() {
		return "", "", "", fmt.Errorf("jwt: token expired: %s", expiresAt.String())
	}

	return userID, login, name, nil
}
