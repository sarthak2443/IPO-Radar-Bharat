package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	Port       string
	LogLevel   string
	DBURL      string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	RedisURL   string
	KafkaURL   string
}

func LoadConfig() (*Config, error) {
	// Attempt to load .env file; ignore if missing in production
	_ = godotenv.Load()

	port := getEnv("HTTP_PORT", "8080")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "ipo_radar")
	dbSSL := getEnv("DB_SSLMODE", "disable")

	dbURL := getEnv("DATABASE_URL", "")
	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser, dbPass, dbHost, dbPort, dbName, dbSSL)
	}

	return &Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		Port:       port,
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		DBURL:      dbURL,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPass,
		DBName:     dbName,
		DBSSLMode:  dbSSL,
		RedisURL:   getEnv("REDIS_URL", "localhost:6379"),
		KafkaURL:   getEnv("KAFKA_BROKERS", "localhost:9092"),
	}, nil
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}
