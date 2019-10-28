package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/rafaelescrich/golden-headed-quetzal/config"
	"github.com/rafaelescrich/golden-headed-quetzal/db"
	"github.com/rafaelescrich/golden-headed-quetzal/files"
)

func migrate() {
	db.DB.AutoMigrate(&files.Metadata{}, &files.Content{})
}

func root(c echo.Context) error {
	return c.JSON(http.StatusOK, "CSV/TXT Uploader and Parser Version 0.0.1")
}

func getFiles(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func getContents(c echo.Context) error {
	contents := "contents"
	return c.JSON(http.StatusOK, contents)
}

func upload(c echo.Context) error {

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	content := make([]byte, file.Size)
	bytesRead, err := src.Read(content)
	if err != nil {
		return c.JSON(403, "Could not read the file: "+err.Error())
	}

	fmt.Printf("%d bytes: %s\n", bytesRead, string(content[:bytesRead]))

	// // Destination
	// dst, err := os.Create(file.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// // Copy
	// if _, err = io.Copy(dst, src); err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusOK, "File "+file.Filename+" uploaded successfully.")
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)

	e.POST("/upload", upload)
	e.GET("/files/:id", getFiles)
	e.GET("/contents", getContents)

	e.Logger.Fatal(e.Start(":1337"))
}
