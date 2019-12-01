package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func DBSetup() {
	db, _ := gorm.Open(
		AppConfig.Database.Type,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			AppConfig.Database.Host,
			AppConfig.Database.Port,
			AppConfig.Database.User,
			AppConfig.Database.DBName,
			AppConfig.Database.Password))
	db.DB().SetMaxIdleConns(10)
	//db.LogMode(true)
	DB = db
	// return DB
}



// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}

