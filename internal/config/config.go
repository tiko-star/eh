package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(k string, d any) any {
	v := os.Getenv(k)
	if v == "" {
		return d
	}
	return v
}
