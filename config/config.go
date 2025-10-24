package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
}

var configuration *Config

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env file", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required.")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service is required.")
		os.Exit(1)
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("http port is required.")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		fmt.Println("Jwt Secret is required.")
		os.Exit(1)
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JwtSecret:   jwtSecret,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
