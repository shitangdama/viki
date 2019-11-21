package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=shitangdama dbname=gaki password=kbr199sd5shi sslmode=disable")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
