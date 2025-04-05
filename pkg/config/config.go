package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	ServerEnv     string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPass        string
	DBName        string
	RedisHost     string
	RedisPort     string
	S3Endpoint    string
	S3AccessKey   string
	S3SecretKey   string
	S3Bucket      string
	S3UseSSL      bool
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		ServerAddress: getRequiredEnv("SERVER_ADDRESS"),
		ServerEnv:     getRequiredEnv("SERVER_ENV"),
		DBHost:        getRequiredEnv("DB_HOST"),
		DBPort:        getRequiredEnv("DB_PORT"),
		DBUser:        getRequiredEnv("DB_USER"),
		DBPass:        getRequiredEnv("DB_PASS"),
		DBName:        getRequiredEnv("DB_NAME"),
		RedisHost:     getRequiredEnv("REDIS_HOST"),
		RedisPort:     getRequiredEnv("REDIS_PORT"),
		S3Endpoint:    getRequiredEnv("S3_ENDPOINT"),
		S3AccessKey:   getRequiredEnv("S3_ACCESS_KEY"),
		S3SecretKey:   getRequiredEnv("S3_SECRET_KEY"),
		S3Bucket:      getRequiredEnv("S3_BUCKET"),
		S3UseSSL:      getRequiredEnvBool("S3_USE_SSL"),
	}

	log.Println("Configuration loaded successfully")
	return cfg
}

func getRequiredEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}

func getRequiredEnvBool(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("Invalid boolean value for %s: %s", key, value)
	}
	return boolValue
}
