package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB()  {
	urlConnection := "host=localhost user=postgres password=root dbname=personalities port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(urlConnection), &gorm.Config{})
	if err != nil {
		log.Panic(err.Error())
	}
}
