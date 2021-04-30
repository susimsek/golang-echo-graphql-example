package config

import "os"

var (
	ServerPort = GetEnv("SERVER_PORT", "9000")
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
