package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(k string, d string) string {
	if os.Getenv(k) != "" {
		return os.Getenv(k)
	} else {
		return d
	}
}

// Config function to get value from env file,
func Config(k string, d string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return GetEnv(k, d)
}
