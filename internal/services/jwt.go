package services

import (
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
		LogErrorf("failed to sign the JWT token for user %q", username)
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
			err := LogAndReturnError("failed to parse JWT token; unexpected signing method")
			return nil, err
		} else {
			key := getSecretKey()
			return key, nil
		}
	})
}

func parseTokenClaims(token *jwt.Token) (string, error) {
	if !token.Valid {
		err := LogAndReturnError("invalid JWT token")
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := LogAndReturnError("failed to parse JWT token; unexpected claims")
		return "", err
	}

	if username, ok := claims["username"].(string); ok {
		return username, nil
	} else {
		err := LogAndReturnError("failed to parse user from JWT token claims")
		return "", err
	}
}
