package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Name string
}

func init() {

}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=shitangdama dbname=gaki password=kbr199sd5shi sslmode=disable")
	if err != nil {
		fmt.Println(err)
		db.Close()
	}
	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu"}
	db.Create(&user)
	result := db.Where("name = ?", "jinzhu").First(&user)
	fmt.Println(result)

	defer db.Close()
}
