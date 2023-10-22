package util

import (
	"os"
)

// GetEnv returns the value of the environment variable with the given key.
// If the environment variable is not set, the default value is returned.
func GetEnv(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}

	return value
}
