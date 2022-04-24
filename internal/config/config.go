package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Host string
	Port string
}

func NewConfig() *Config {
	return &Config{
		Env: getenv("APP_ENV", "local"),
	}
}

func (c *Config) Init() {
	c.LoadEnv()

	c.Host = getenv("APP_HOST", "localhost")
	c.Port = getenv("APP_PORT", "8080")
}

func (c *Config) LoadEnv() {
	filename := ".env." + c.Env

	err := godotenv.Load("../env/" + filename)
	if err != nil {
		log.Fatalf("Error loading %s file", filename)
	}

	log.Printf("Loaded environment variables from %s", filename)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("%s environment variable not set. Using fallback: %s", key, fallback)
		return fallback
	}
	return value
}
