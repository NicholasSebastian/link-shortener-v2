package main

import (
	"log"

	"github.com/NicholasSebastian/link-shortener-v2/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Server Error: Failed to load environment variables.")
	}

	username := ""
	tokenstr := services.NewJwtToken(username)
	// TODO: Store this in the user's cookies or payload or sth.
	// TODO: All the other stuff...

	log.Println(tokenstr)
}
