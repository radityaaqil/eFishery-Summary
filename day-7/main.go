package main

import (
	"day7/connection"
	"day7/router"
	"fmt"
	"log"
)

func main() {

	config := connection.ConfigurationDB()

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
