package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ConvertID(id string) (uint, error) {
	newID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(newID), nil
}
