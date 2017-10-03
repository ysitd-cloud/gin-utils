package env

import "os"

func GetEnvWithDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
