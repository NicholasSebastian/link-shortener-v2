package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NicholasSebastian/link-shortener-v2/internal/services"
	"github.com/joho/godotenv"
)

const PORT = 8000

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load environment variables")
	}

	mux := http.NewServeMux()
	// TODO: Register the route handlers here.

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error running the HTTP server: %s\n", err)
	}

	username := ""
	tokenstr := services.NewJwtToken(username)
	// TODO: Store this in the user's cookies or payload or sth.
	// TODO: Move this into the API handler for '/login'

	// TODO: All the other stuff...

	log.Println(tokenstr)
}
