package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	JwtSecret string
}

var instance *Config = nil

func (c *Config) GetJWTSecret() string {
	return c.JwtSecret
}

func GetConfig() (*Config, error) {
	if instance == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		instance = &Config{JwtSecret: os.Getenv("JWT_SECRET")}
	}
	return instance, nil
}
