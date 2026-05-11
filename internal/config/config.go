package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string
	DSN     string
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("key: \"%s\" does not exist inside .env", key))
}

func Load() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT"),
		DSN: fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			getEnv("DB_HOST"),
			getEnv("DB_PORT"),
			getEnv("DB_USER"),
			getEnv("DB_PASSWORD"),
			getEnv("DB_NAME"),
			getEnv("DB_SSLMODE"),
		),
	}
}
