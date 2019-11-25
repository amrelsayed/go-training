package main

import (
	"fmt"

	"github.com/amrelsayed/training/todos/config"
	models "github.com/amrelsayed/training/todos/models"
	"github.com/amrelsayed/training/todos/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {

	// Creating a connection to the database
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run(":8081")
}
