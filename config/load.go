package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Uri      string
	Broker   string
	User     string
	Password string
}

func Load() Config {
	_ = godotenv.Load()
	uri := os.Getenv("mongo-uri")
	broker := os.Getenv("broker")
	user := os.Getenv("user")
	password := os.Getenv("password")

	config := Config{
		Uri:      uri,
		Broker:   broker,
		User:     user,
		Password: password,
	}

	fmt.Println(config)

	return config
}
