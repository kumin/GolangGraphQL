package envx

import (
	"os"
	"strconv"
)

func GetString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultVal
	}

	return value
}

func GetInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	value, err := strconv.ParseInt(valStr, 10, 32)
	if err != nil {
		return defaultVal
	}

	return int(value)
}
