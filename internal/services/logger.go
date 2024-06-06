package services

import (
	"fmt"
	"log"
)

func LogError(message string) {
	log.Println(message)

	// TODO: Save to the database.
}

func LogAndReturnError(message string) error {
	LogError(message)
	return fmt.Errorf(message)
}
