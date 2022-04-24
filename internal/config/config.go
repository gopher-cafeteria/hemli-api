package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	c.LoadEnv()

	c.Host = os.Getenv("APP_HOST")
	c.Port = os.Getenv("APP_PORT")
}

func (c *Config) LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	filename := ".env." + env
	path := "../env/" + filename

	if env == "production" {
		path = "/etc/secrets/" + filename
	}

	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading %s file", filename)
	}

	log.Printf("Loaded environment variables from %s", filename)
}
