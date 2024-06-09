package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NicholasSebastian/link-shortener-v2/internal/auth"
	"github.com/joho/godotenv"
)

const PORT = 8000

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load environment variables")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.HandleLogin)
	// TODO: Register the route handlers here.

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error running the HTTP server: %s\n", err)
	}

	// TODO: All the other stuff...
}
