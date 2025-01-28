package utils

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

const (
	alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length       = 32
)

func GenerateId() string {
	id, err := gonanoid.Generate(alphaNumeric, length)
	if err != nil {
		log.Println("Error generating ID =>", err)
		return ""
	}

	return id
}
