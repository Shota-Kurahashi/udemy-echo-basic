package router

import (
	"ge-rest-api/src/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	router := echo.New()

	router.POST("/signup", uc.SignUp)
	router.POST("/login", uc.LogIn)
	router.POST("/logout", uc.LogOut)

	return router
}
