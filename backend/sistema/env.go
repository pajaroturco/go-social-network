package sistema

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(llave string) string {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	return os.Getenv(llave)
}