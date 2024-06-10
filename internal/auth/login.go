package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NicholasSebastian/link-shortener-v2/internal/services"
)

type credentials struct {
	Username string
	Password string
}

func HandleLogin(res http.ResponseWriter, req *http.Request) {
	var c credentials
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&c); err != nil {
		services.LogErrorf("Malformed Login Request Body: %s", err)
	}

	// TODO: Check if the given username and password are valid.
	log.Println(c.Username, c.Password)

	tokenstr := NewJwtToken(c.Username)
	log.Println(tokenstr)
	// TODO: Store this in the user's cookies or payload or sth.

	// TODO: Respond with the appropriate status code.
}
