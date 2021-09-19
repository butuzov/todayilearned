package testecho

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secret = []byte("secret")

type TokenClaim struct {
	*jwt.StandardClaims
	TokenType string
	Name      string
}

func verifyTokenWithKey(token *jwt.Token) (interface{}, error) {
	return secret, nil
}

func createToken(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// set our claims
	token.Claims = &TokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
		"root",
		user,
	}

	return token.SignedString(secret)
}
