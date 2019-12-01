package main

import (
	"log"
	"fmt"
	"net/http"
	"viki/controllers"
	"viki/common"
	"viki/model"
	"viki/pkg/logging"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	common.ConfigSetup()
	common.DBSetup()
	model.Migration()
	logging.Setup()
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


	log.Printf("[info] start http server listening %s", ":8080")
	s.ListenAndServe()
}
