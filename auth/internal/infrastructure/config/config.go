package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App      App
	Postgres Postgres
}

type App struct {
	Name string
	Port int
}

type Postgres struct {
	Database string
	User     string
	Password string
	Port     int
}

func New() (Config, error) {
	appName, err := read("APP_NAME")
	if err != nil {
		return Config{}, fmt.Errorf("reading APP_NAME: %w", err)
	}

	appPortRaw, err := read("APP_PORT")
	if err != nil {
		return Config{}, fmt.Errorf("reading APP_PORT: %w", err)
	}

	appPort, err := strconv.Atoi(appPortRaw)
	if err != nil {
		return Config{}, fmt.Errorf("converting APP_PORT to int: %w", err)
	}

	postgresDatabase, err := read("POSTGRES_DB")
	if err != nil {
		return Config{}, fmt.Errorf("reading POSTGRES_DB: %w", err)
	}

	postgresUser, err := read("POSTGRES_USER")
	if err != nil {
		return Config{}, fmt.Errorf("reading POSTGRES_USER: %w", err)
	}

	postgresPassword, err := read("POSTGRES_PASSWORD")
	if err != nil {
		return Config{}, fmt.Errorf("reading POSTGRES_PASSWORD: %w", err)
	}

	postgresPortRaw, err := read("POSTGRES_PORT")
	if err != nil {
		return Config{}, fmt.Errorf("reading POSTGRES_PORT: %w", err)
	}

	postgresPort, err := strconv.Atoi(postgresPortRaw)
	if err != nil {
		return Config{}, fmt.Errorf("converting POSTGRES_PORT to int: %w", err)
	}

	return Config{
		App: App{
			Name: appName,
			Port: appPort,
		},
		Postgres: Postgres{
			Database: postgresDatabase,
			User:     postgresUser,
			Password: postgresPassword,
			Port:     postgresPort,
		},
	}, nil
}

func read(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("the variable is required")
	}

	return value, nil
}
