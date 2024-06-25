package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

type Config struct {
	PostgresURI     string
	MaxConnections  int32
	MinConnections  int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
	Host            string
	Port            string
}

func NewConfig() *Config {
	_ = godotenv.Load("./build/.env")
	postgresURI := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	maxConnections, err := strconv.Atoi(os.Getenv("MAX_CONNECTIONS"))
	minConnections, err := strconv.Atoi(os.Getenv("MIN_CONNECTIONS"))
	maxConnLifetime, err := strconv.Atoi(os.Getenv("MAX_CONN_LIFE_TIME"))
	MaxConnIdleTime, err := strconv.Atoi(os.Getenv("MAX_CONN_IDLE_TIME"))
	if err != nil {
		log.Fatalf("Could not parse variables: %v", err.Error())
	}
	return &Config{
		PostgresURI:     postgresURI,
		MaxConnections:  int32(maxConnections),
		MinConnections:  int32(minConnections),
		MaxConnLifetime: time.Duration(maxConnLifetime) * time.Millisecond,
		MaxConnIdleTime: time.Duration(MaxConnIdleTime) * time.Millisecond,
		Host:            os.Getenv("HOST"),
		Port:            os.Getenv("PORT"),
	}
}
