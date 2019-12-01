package main

import (
	"fmt"
	"net/http"
	"viki/controllers"
	"viki/common"
	"viki/model"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	common.ConfigSetup()
	common.DBSetup()
	model.Migration()
}

func main() {
	router := controllers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: router,
		// ReadTimeout:    setting.ReadTimeout,
		// WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
