package services

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		tokenstr := "" // TODO: Get the token out of the request cookies/header.

		token, err := parseTokenMethod(tokenstr)
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			// TODO: redirect to login page.
		}

		username, err := parseTokenClaims(token)
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			// TODO: redirect to login page.
		}

		log.Printf("successfully logged in as user %q", username)
		next(res, req)
	}
}
