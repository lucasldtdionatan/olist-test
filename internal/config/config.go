package config

import "os"

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() Config {
	return Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
	}
}
