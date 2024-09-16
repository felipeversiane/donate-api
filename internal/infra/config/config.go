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
	Database     DatabaseConfig
	Server       ServerConfig
	Log          LogConfig
	CloudService CloudServiceConfig
}

type ConfigInterface interface {
	GetDatabaseConfig() DatabaseConfig
	GetServerConfig() ServerConfig
	GetLogConfig() LogConfig
	GetCloudServiceConfig() CloudServiceConfig
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

type LogConfig struct {
	Level string
}

type CloudServiceConfig struct {
	Region          string
	AccessKey       string
	SecretAccessKey string
	Endpoint        string
	ObjectStorage   ObjectStorageConfig
}

type ObjectStorageConfig struct {
	Bucket string
	URL    string
	ACL    string
}

func NewConfig() ConfigInterface {
	var cfg *Config
	once.Do(func() {
		cfg = &Config{
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
			Log: LogConfig{
				Level: getEnvOrDie("LOG_LEVEL"),
			},
			CloudService: CloudServiceConfig{
				Region:          getEnvOrDie("AWS_REGION"),
				AccessKey:       getEnvOrDie("AWS_ACCESS_KEY_ID"),
				SecretAccessKey: getEnvOrDie("AWS_SECRET_ACCESS_KEY"),
				Endpoint:        getEnvOrDie("AWS_ENDPOINT"),
				ObjectStorage: ObjectStorageConfig{
					Bucket: getEnvOrDie("S3_BUCKET"),
					URL:    getEnvOrDie("S3_URL"),
					ACL:    getEnvOrDie("S3_ACL"),
				},
			},
		}
	})
	return cfg
}

func getEnvOrDie(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Errorf("missing environment variable %s", key))
	}
	return value
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *Config) GetServerConfig() ServerConfig {
	return c.Server
}

func (c *Config) GetLogConfig() LogConfig {
	return c.Log
}

func (c *Config) GetCloudServiceConfig() CloudServiceConfig {
	return c.CloudService
}
