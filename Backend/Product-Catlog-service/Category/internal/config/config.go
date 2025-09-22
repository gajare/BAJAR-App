package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() (*Config, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=postgres dbname=product_catalog port=5432 sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	return &Config{
		DatabaseURL: databaseURL,
		Port:        port,
	}, nil
}
