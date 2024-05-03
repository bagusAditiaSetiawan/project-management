package config

import (
	"github.com/bagusAditiaSetiawan/project-management/api/helpers"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	helpers.IfPanicHelper(err)
	return os.Getenv(key)
}
