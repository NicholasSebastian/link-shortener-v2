package services

import (
	"fmt"
	"log"
)

func LogError(message string) {
	log.Println(message)

	// TODO: Save to the database.
}

func LogErrorf(format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	LogError(message)
}

func LogAndReturnError(message string) error {
	LogError(message)
	return fmt.Errorf(message)
}
