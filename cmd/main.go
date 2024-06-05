package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Server Error: Failed to load environment variables.")
	}

	// TODO: man what am i doing with my life.

	authKey := os.Getenv("AUTH_KEY")
	authKeyBytes := []byte(authKey)

	log.Println(authKeyBytes) // just testing
}
