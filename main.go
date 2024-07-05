package main

import (
	"fmt"

	"github.com/askrishna98/library_/handlers"
	"github.com/askrishna98/library_/loaddata"
	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
)

func main() {
	// should initiate DB first
	DB := models.GetMockDBInstance()
	Id := service.InitalizeIDGenerator()

	loaddata.LoadData(DB, Id)

	for _, val := range DB.Books {
		fmt.Println(*val)
	}
	for _, val := range DB.Members {
		fmt.Println(*val)
	}
	handlers.StartApp()

}
