package main

import (
	"ge-rest-api/src/controller"
	"ge-rest-api/src/db"
	"ge-rest-api/src/repository"
	"ge-rest-api/src/router"
	"ge-rest-api/src/usecase"
)

func main() {
	db := db.Connect()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUsecase)

	router := router.NewRouter(userController)

	router.Logger.Fatal(router.Start(":8080"))
}
