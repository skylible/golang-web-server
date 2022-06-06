package main

import (
	"fmt"
	"web-server/config"
	"web-server/models"
	"web-server/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	r := routes.SetupRouter()
	//running
	r.Run("localhost:3000")
}
