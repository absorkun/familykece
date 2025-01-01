package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initdb() {
	var err error
	var dsn string = os.Getenv("DATABASE_URL")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&Product{}, &Category{})
	var category Category
	if err := db.FirstOrCreate(&category, Category{
		ID:   "COMP",
		Name: "Computer",
	}).Error; err != nil {
		log.Fatal(err.Error())
	}

	var product Product
	if err := db.FirstOrCreate(&product, Product{
		ID:         "LEN001",
		Name:       "Thinpad X280",
		Price:      10000000,
		CategoryID: "COMP",
		Icon:       "len001.png",
	}).Error; err != nil {
		log.Fatal(err.Error())
	}
}
