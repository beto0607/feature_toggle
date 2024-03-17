package configs

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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

func EnvHostname() string {
	LoadDotEnv()

	hostname, found := os.LookupEnv("HOSTNAME")
	if found {
		return hostname
	}
	return "localhost"
}

func ShouldConnectDB() bool {
	return true 
}

func DefaultReadTimeout() time.Duration {
	return 10 * time.Second
}
func DefaultWriteTimeout() time.Duration {
	return 10 * time.Second
}
