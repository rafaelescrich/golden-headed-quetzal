package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func root(c echo.Context) error {
	return c.JSON(http.StatusOK, "CSV/TXT Uploader and Parser Version 0.0.1")
}

func getFiles(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
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

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "File "+file.Filename+" uploaded successfully.")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)

	e.POST("/upload", upload)
	e.GET("/files/:id", getFiles)

	e.Logger.Fatal(e.Start(":1337"))
}
