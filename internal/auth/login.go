package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NicholasSebastian/link-shortener-v2/internal/models"
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
		services.LogErrorf("malformed login request body: %s", err)
	}

	database := services.NewDatabase()
	users := models.NewUsers(database)

	valid, err := users.Exists(c.Username, c.Password)
	if err != nil {
		services.LogErrorf("failed to access users database: %s", err)
	}

	if valid {
		tokenstr := NewJwtToken(c.Username)
		log.Println(tokenstr)
		// TODO: Store this in the user's cookies or payload or sth.
	} else {
		res.WriteHeader(http.StatusUnauthorized)
	}
}
