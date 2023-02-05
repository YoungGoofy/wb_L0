package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	MaxAttempts      string
}

func (config *Config) InitFromEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	config.DatabaseUser = os.Getenv("dbuser")
	config.DatabasePassword = os.Getenv("dbpasswd")
	config.DatabaseHost = os.Getenv("dbhost")
	config.DatabasePort = os.Getenv("dbport")
	config.DatabaseName = os.Getenv("dbname")
	config.MaxAttempts = os.Getenv("maxattempts")
	return nil
}
