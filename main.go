package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/routes"
)

func main() {
	//connect DB
	db := config.Connect()
	defer db.Close()

	// migrate DB
	config.MigrateDB()

	// routes controller
	controller := routes.Controller()

	//log
	fmt.Println("Application Running in Port : 8080")
	log.Fatal(http.ListenAndServe(":8080", controller))
}
