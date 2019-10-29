package main

import (
	"log"

	"github.com/labstack/echo"

	"github.com/rafaelescrich/golden-headed-quetzal/config"
	"github.com/rafaelescrich/golden-headed-quetzal/db"
	"github.com/rafaelescrich/golden-headed-quetzal/files"
	"github.com/rafaelescrich/golden-headed-quetzal/router"
)

func migrate() {
	db.DB.AutoMigrate(&files.Metadata{}, &files.Content{})
}

func main() {

	err := config.Load()

	if err != nil {
		log.Fatal("Error while initializing config: ", err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}

	// Add tables to db if they dont exist
	migrate()

	e := echo.New()

	// Add routes and handlers
	router.NewRouter(e)

	e.Logger.Fatal(e.Start(":1337"))
}
