package db

import (
	"github.com/jinzhu/gorm"
	"github.com/rafaelescrich/golden-headed-quetzal/config"

	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB instantiate a global variable with the database connection
var DB *gorm.DB

// Connect to postgres database
func Connect() (err error) {

	cfg := config.Cfg.Database

	dbConfig := "host=" + cfg.Host + " port=" + cfg.Port + " user=" + cfg.User + " password=" + cfg.Password + " sslmode=disable"

	DB, err = gorm.Open("postgres", dbConfig)
	return
}
