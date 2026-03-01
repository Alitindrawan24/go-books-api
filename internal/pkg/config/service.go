package config

import (
	"os"
	"strconv"
)

// Get retrieves environment variable value by key
// Returns empty string if not found
func Get(key string) string {
	return os.Getenv(key)
}

// GetWithDefault retrieves environment variable or returns default if empty
func GetWIthDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

// GetInt64WithDefault retrieves environment variable as int64 or returns default if empty
func GetInt64WithDefault(key string, defaultValue int64) int64 {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	valueInt, _ := strconv.Atoi(value)

	return int64(valueInt)
}

// GetFloat64WithDefault retrieves environment variable as float64 or returns default if empty
func GetFloat64WithDefault(key string, defaultValue float64) float64 {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	valueFloat, _ := strconv.ParseFloat(value, 64)

	return float64(valueFloat)
}
