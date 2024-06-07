package utils

import (
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func LoadEnvVars(keys ...string) map[string]string{
	envVars := make(map[string]string)

	if err:= godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	for _, key := range keys {
		envVars[key] = os.Getenv(key)
	}

	return envVars
}