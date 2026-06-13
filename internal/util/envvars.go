package util

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func MustGetVariable(key string) (value string) {
	value = os.Getenv(strings.ToUpper(key))
	if value == "" {
		log.Fatalf("Couldn't find value for key %s", key)
	}
	return value
}

func GetEnvVariable(key, fallback string) (value string) {
	value = os.Getenv(strings.ToUpper(key))
	if value == "" {
		fmt.Printf("Couldn't find value for key %s, using fallback %s\n", key, fallback)
		return fallback
	}
	return value
}
