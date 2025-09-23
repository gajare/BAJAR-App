package config

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() Config {
	return Config{
		DBUrl: "host=localhost user=postgres password=1234 dbname=paymentsdb port=5432 sslmode=disable",
		Port:  "8080",
	}
}
