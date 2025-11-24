package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerPort string
	LogLevel   string

	RabbitExchange   string
	RabbitURI        string
	RabbitRoutingKey string

	MinioEndpoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioUseSSL    bool
	MinioBucket    string
}

func LoadConfig() *Config {
	slog.Info("Loading config...")

	if err := godotenv.Load(".env"); err != nil {
		slog.Warn(".env file was not used in configuration")
	}
	// Generic
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}

	// RabbitMQ
	rabbitExchange := os.Getenv("RABBITMQ_EXCHANGE")
	if rabbitExchange == "" {
		slog.Error("No RABBITMQ_EXCHANGE specified")
	}
	rabbitURI := os.Getenv("RABBITMQ_URI")
	if rabbitURI == "" {
		slog.Error("No RABBITMQ_URI specified")
	}
	rabbitRoutingKey := os.Getenv("RABBITMQ_ROUTING_KEY")
	if rabbitRoutingKey == "" {
		slog.Error("No RABBITMQ_ROUTING_KEY specified")
	}

	// Minio
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	if minioEndpoint == "" {
		slog.Error("No MINIO_ENDPOINT specified")
	}
	minioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	if minioAccessKey == "" {
		slog.Error("No MINIO_ACCESS_KEY specified")
	}
	minioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	if minioSecretKey == "" {
		slog.Error("No MINIO_SECRET_KEY specified")
	}

	useSSL := os.Getenv("MINIO_USE_SSL")
	if useSSL == "" {
		slog.Error("No MINIO_USE_SSL specified")
	}
	minioSSL, _ := strconv.ParseBool(useSSL)

	minioBucket := os.Getenv("MINIO_BUCKET")
	if minioBucket == "" {
		slog.Error("No MINIO_BUCKET specified")
	}

	slog.Info("Loading config successful")
	return &Config{
		ServerPort:       serverPort,
		LogLevel:         logLevel,
		RabbitExchange:   rabbitExchange,
		RabbitURI:        rabbitURI,
		RabbitRoutingKey: rabbitRoutingKey,
		MinioEndpoint:    minioEndpoint,
		MinioAccessKey:   minioAccessKey,
		MinioSecretKey:   minioSecretKey,
		MinioUseSSL:      minioSSL,
		MinioBucket:      minioBucket,
	}
}
