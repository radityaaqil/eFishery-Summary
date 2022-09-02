package main

import (
	"day9/connection"
	"day9/router"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	//Dotenv
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &connection.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	//Database Connection
	db, err := connection.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected", db)

	//Start Server
	e := router.InitRoute(db)
	e.Logger.Fatal(e.Start("127.0.0.1:5000"))
}
