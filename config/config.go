package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecret   string
	DB          *DBConfig
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

	host := os.Getenv("DB_HOST")
	if host == "" {
		fmt.Println("DB Host is required.")
		os.Exit(1)
	}
	dbPortStr := os.Getenv("DB_PORT")
	if dbPortStr == "" {
		fmt.Println("DB Port is required.")
		os.Exit(1)
	}
	dbPort, err := strconv.ParseInt(dbPortStr, 10, 64)
	if err != nil {
		fmt.Println("DB Port must be number")
		os.Exit(1)
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		fmt.Println("DB User is required.")
		os.Exit(1)
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		fmt.Println("DB Password is required.")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required.")
		os.Exit(1)
	}
	sslMode := os.Getenv("DB_SSLMODE")
	if sslMode == "" {
		fmt.Println("DB SSLMode is required.")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:     host,
		Port:     int(dbPort),
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JwtSecret:   jwtSecret,
		DB:          dbConfig,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
