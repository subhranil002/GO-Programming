package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port            int
	AppEnv          string
	MongoURI        string
	MongoDatabase   string
	MongoCollection string
}

func Load() (Config, error) {
	port, err := getPort()
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Port:            port,
		AppEnv:          getEnv("APP_ENV", "development"),
		MongoURI:        strings.TrimSpace(os.Getenv("MONGODB_URI")),
		MongoDatabase:   strings.TrimSpace(os.Getenv("MONGODB_DATABASE")),
		MongoCollection: strings.TrimSpace(os.Getenv("MONGODB_COLLECTION")),
	}

	switch {
	case cfg.MongoURI == "":
		return Config{}, errors.New("MONGODB_URI is required")

	case cfg.MongoDatabase == "":
		return Config{}, errors.New("MONGODB_DATABASE is required")

	case cfg.MongoCollection == "":
		return Config{}, errors.New("MONGODB_COLLECTION is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))

	if value == "" {
		return fallback
	}

	return value
}

func getPort() (int, error) {
	value := strings.TrimSpace(os.Getenv("PORT"))

	if value == "" {
		return 3000, nil
	}

	port, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("PORT must be an integer: %w", err)
	}

	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("PORT must be between 1 and 65535")
	}

	return port, nil
}