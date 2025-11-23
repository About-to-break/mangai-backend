package config

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	ServerPort     string
	LogLevel       string
	RabbitExchange string
	RabbitURI      string
}

func LoadConfig() *Config {
	slog.Info("Loading config...")

	if err := godotenv.Load(".env"); err != nil {
		slog.Warn(".env file was not used in configuration")
	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}
	rabbitExchange := os.Getenv("RABBITMQ_EXCHANGE")
	if rabbitExchange == "" {
		slog.Error("No RABBITMQ_EXCHANGE specified")
	}
	rabbitURI := os.Getenv("RABBITMQ_URI")
	if rabbitURI == "" {
		slog.Error("No RABBITMQ_URI specified")
	}
	slog.Info("Loading config successful")
	return &Config{
		ServerPort:     serverPort,
		LogLevel:       logLevel,
		RabbitExchange: rabbitExchange,
		RabbitURI:      rabbitURI,
	}
}
