package tools

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SERVER_KEY_DEV = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3PqZEDzJ2E8le5aFs8Tw
Um0tcUrc+614d9fseI6pmVOTKcNWTgktNX9rTz/B4JTCws3/8erqMVkwuz1vhH6S
iY+BUyn24g44/szZtVc0RgZVqLnZ87nsWvL2C+M1L4AiIgAwyElOFY5MCuknXMxD
oYOmobEYbJry4+ZkraUUzCGgWIDoYK8j/JzG63mw6QUZPO9fKSgUPDyjh6NAOq2c
+4WcdG2Ss0mseYeUXjUW0S2IZTXkYcfqJyXjDgCNSUGUzA5+NwKyjl5Sijr55ULD
wUMqJYfjFEGN4HlIqA80PHpQqjiWtAekZNRbfu0yhUW1s1ZUQchw1R7LF28aq5Bo
8wIDAQAB
-----END PUBLIC KEY-----`
)

func VerifyJWT(jwtToken string) (string, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(SERVER_KEY_DEV))
	if err != nil {
		return "", err
	}
	var publicKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return publicKey, nil }

	token, err := jwt.Parse(jwtToken, publicKeyFunc)
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	exp := claims["exp"].(float64)
	subject := claims["sub"].(string)

	expiresAt := time.Unix(int64(exp), 0)

	if int64(exp) < time.Now().Unix() {
		return "", fmt.Errorf("JWT Token expired: %s", expiresAt.String())
	}

	return subject, nil
}
