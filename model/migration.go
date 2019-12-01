package model

import (
	"viki/common"
)

func Migration() {
	db := common.GetDB()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
}