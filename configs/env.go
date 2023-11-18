package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var dotEnvLoaded = false

func LoadDotEnv() {
	if dotEnvLoaded {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	dotEnvLoaded = true
}

func EnvMongoURI() string {
	LoadDotEnv()

	return os.Getenv("MONGODB_URI")
}

func EnvPort() string {
	LoadDotEnv()

	port, found := os.LookupEnv("PORT")
	if found {
		return port
	}
	return "8080"

}
