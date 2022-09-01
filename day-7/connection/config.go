package connection

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConfigurationDB() *Config {
	//Dotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	return config
}
