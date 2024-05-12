package cmd

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr     string
	SupertokensURI string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadConfig() Config {
	godotenv.Load()

	// default setting
	cfg := Config{
		ServerAddr:     "localhost:2000",
		SupertokensURI: "http://localhost:3567",
		DBHost:         "localhost",
		DBPort:         "5432",
		DBUser:         "postgres",
		DBPass:         "postgres",
		DBName:         "postgres",
	}

	if supertokensUri, ok := os.LookupEnv("SUPERTOKENS_URI"); ok && supertokensUri != "" {
		cfg.SupertokensURI = supertokensUri
	}

	if serverAddr, ok := os.LookupEnv("SERVER_ADDR"); ok && serverAddr != "" {
		cfg.ServerAddr = serverAddr
	}

	if dbHost, ok := os.LookupEnv("DB_HOST"); ok && dbHost != "" {
		cfg.DBHost = dbHost
	}

	if dbPort, ok := os.LookupEnv("DB_PORT"); ok && dbPort != "" {
		cfg.DBPort = dbPort
	}

	if dbUser, ok := os.LookupEnv("DB_USER"); ok && dbUser != "" {
		cfg.DBUser = dbUser
	}

	if dbPass, ok := os.LookupEnv("DB_PASS"); ok && dbPass != "" {
		cfg.DBPass = dbPass
	}

	if dbName, ok := os.LookupEnv("DB_NAME"); ok && dbName != "" {
		cfg.DBName = dbName
	}

	return cfg
}
