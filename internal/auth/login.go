package auth

import (
	"log"
	"net/http"
)

func HandleLogin(res http.ResponseWriter, req *http.Request) {
	username := ""
	password := "" // TODO: Retrieve these from the request body.

	// TODO: Check if the given username and password are valid.
	log.Println(username, password)

	tokenstr := NewJwtToken(username)
	log.Println(tokenstr)
	// TODO: Store this in the user's cookies or payload or sth.

	// TODO: Respond with the appropriate status code.
}
