package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	urls = []string{"https://habr.com/ru/rss/hub/go/all/?fl=ru", "https://habr.com/ru/rss/best/daily/?fl=ru", "https://cprss.s3.amazonaws.com/golangweekly.com.xml"}
)

type Config struct {
	ServiceName string
	DatabaseURL string
	ServerHost  string
	UpdateTime  string
	Urls        []string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ServiceName: os.Getenv("SERVICE_NAME"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		ServerHost:  os.Getenv("SERVER_HOST"),
		UpdateTime:  os.Getenv("UPDATE_TIME"),
		Urls:        urls,
	}, nil
}
