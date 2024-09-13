package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	once sync.Once
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

func NewConfig() *Config {
	var config *Config
	once.Do(func() {
		config = &Config{
			Database: DatabaseConfig{
				Host:     getEnvOrDie("POSTGRES_HOST"),
				Port:     getEnvOrDie("POSTGRES_PORT"),
				User:     getEnvOrDie("POSTGRES_USER"),
				Password: getEnvOrDie("POSTGRES_PASSWORD"),
				Name:     getEnvOrDie("POSTGRES_DB"),
			},
			Server: ServerConfig{
				Port: getEnvOrDie("API_PORT"),
			},
		}
	})
	return config
}

func getEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("missing environment variable %s", key))
	}
	return value
}
