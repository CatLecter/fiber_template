// Package config содержит конфигурацию приложения.
package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config содержит все настройки приложения.
type Config struct {
	Port            string
	PostgresURI     string
	MaxConnections  int32
	MinConnections  int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

// NewConfig создает новый экземпляр конфигурации.
func NewConfig() *Config {
	if err := godotenv.Load("configs/.env"); err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Собираем строку подключения к PostgreSQL из отдельных переменных
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresDB := os.Getenv("POSTGRES_DB")

	if postgresUser == "" || postgresPassword == "" || postgresHost == "" || postgresPort == "" || postgresDB == "" {
		panic("PostgreSQL configuration is incomplete")
	}

	hostPort := net.JoinHostPort(postgresHost, postgresPort)
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", postgresUser, postgresPassword, hostPort, postgresDB)

	maxConnections, err := strconv.ParseInt(os.Getenv("MAX_CONNECTIONS"), 10, 32)
	if err != nil {
		maxConnections = 10
	}

	minConnections, err := strconv.ParseInt(os.Getenv("MIN_CONNECTIONS"), 10, 32)
	if err != nil {
		minConnections = 2
	}

	maxConnLifetime, err := time.ParseDuration(os.Getenv("MAX_CONN_LIFETIME"))
	if err != nil {
		maxConnLifetime = 5 * time.Minute
	}

	maxConnIdleTime, err := time.ParseDuration(os.Getenv("MAX_CONN_IDLE_TIME"))
	if err != nil {
		maxConnIdleTime = 5 * time.Minute
	}

	return &Config{
		Port:            port,
		PostgresURI:     postgresURI,
		MaxConnections:  int32(maxConnections),
		MinConnections:  int32(minConnections),
		MaxConnLifetime: maxConnLifetime,
		MaxConnIdleTime: maxConnIdleTime,
	}
}
