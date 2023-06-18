package main

import (
	"fmt"
	"ge-rest-api/src/db"
	"ge-rest-api/src/model"
)

func main() {
	database := db.Connect()
	defer db.Close(database)
	defer fmt.Println("Database connection closed")
	database.AutoMigrate(&model.User{}, &model.Task{})
}
