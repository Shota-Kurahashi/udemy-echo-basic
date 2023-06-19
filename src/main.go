package main

import (
	"ge-rest-api/src/controller"
	"ge-rest-api/src/db"
	"ge-rest-api/src/repository"
	"ge-rest-api/src/router"
	"ge-rest-api/src/usecase"
	"ge-rest-api/src/validator"
)

func main() {
	db := db.Connect()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	router := router.NewRouter(userController, taskController)

	router.Logger.Fatal(router.Start(":8080"))
}
