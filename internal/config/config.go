package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB
	HTTPServer
}

type DB struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     int64  `env:"DB_PORT"`
}

type HTTPServer struct {
	URL string `env:"SERVER_URL" env-default:"localhost:7070"`
}

func ParseConfig() (*Config, error) {
	var config Config
	err := cleanenv.ReadConfig("./config/.env", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
