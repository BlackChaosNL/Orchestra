package config

import (
	"os"
	"strconv"
)

func GetStrFromEnv(k string, d string) string {
	if os.Getenv(k) != "" {
		return os.Getenv(k)
	}
	return d
}

func GetBoolFromEnv(k string, d bool) bool {
	if os.Getenv(k) != "" {
		bool, _ := strconv.ParseBool(os.Getenv(k))
		return bool
	}
	return d
}
