package utilities

import "os"

func GetEnv(envName, defaultValue string) string {
	value := defaultValue
	if envValue := os.Getenv(envName); envValue != "" {
		value = envValue
	}
	return value
}
