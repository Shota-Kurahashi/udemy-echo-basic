package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetDbUrl() string {

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", GetEnv("POSTGRES_USER"), GetEnv("POSTGRES_PW"), GetEnv("POSTGRES_HOST"), GetEnv("POSTGRES_PORT"), GetEnv("POSTGRES_DB"))
}

func GetEnv(key string) string {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	return os.Getenv(key)
}

func CheckError(err error, args ...interface {
}) {
	if err != nil {
		log.Fatalln(err, args)
	}
}
