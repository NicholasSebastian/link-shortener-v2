package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SESSION_MINS = 30

func NewJwtToken(username string) string {
	claims := jwt.MapClaims{
		"username": username,
		"expire":   time.Now().Add(time.Minute * SESSION_MINS),
	}

	key := getSecretKey()
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	tokenstr, err := token.SignedString(key)

	if err != nil {
		msg := fmt.Sprintf("failed to sign the JWT token for user %q", username)
		LogError(msg)
	}
	return tokenstr
}

func getSecretKey() []byte {
	secretKey := os.Getenv("AUTH_KEY")
	return []byte(secretKey)
}

func parseTokenMethod(tokenstr string) (*jwt.Token, error) {
	return jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			msg := "failed to parse JWT token; unexpected signing method"
			err := LogAndReturnError(msg)
			return nil, err
		} else {
			key := getSecretKey()
			return key, nil
		}
	})
}

func parseTokenClaims(token *jwt.Token) (string, error) {
	if !token.Valid {
		msg := "invalid JWT token"
		err := LogAndReturnError(msg)
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		msg := "failed to parse JWT token; unexpected claims"
		err := LogAndReturnError(msg)
		return "", err
	}

	if username, ok := claims["username"].(string); ok {
		return username, nil
	} else {
		msg := "failed to parse username from JWT token claims; expected a string"
		err := LogAndReturnError(msg)
		return "", err
	}
}
