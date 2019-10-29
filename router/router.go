package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rafaelescrich/golden-headed-quetzal/files"
)

// NewRouter creates routes and add middlewares
func NewRouter(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", root)

	e.POST("/upload", upload)
	e.GET("/files/:id", getFiles)
	e.GET("/contents", getContents)
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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(403, "Could not read the file: "+err.Error())
	}
	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(403, "Could not read the file: "+err.Error())
	}
	defer file.Close()

	err = files.Save(fileHeader.Filename, fileHeader.Size, file)
	if err != nil {
		return c.JSON(403, "Could not save on database: "+err.Error())
	}

	return c.JSON(http.StatusOK, "File "+fileHeader.Filename+" uploaded successfully.")
}
