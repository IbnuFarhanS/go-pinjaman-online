package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string, v ...any) string {
	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	if key != "" {
		return os.Getenv(key)
	}
	return v[0].(string)
}
