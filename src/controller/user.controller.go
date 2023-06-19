package controller

import (
	"ge-rest-api/src/model"
	"ge-rest-api/src/usecase"
	"ge-rest-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type useController struct {
	uu usecase.IUserUseCase
}

func NewUserController(uu usecase.IUserUseCase) IUserController {
	return &useController{uu}
}

func (uc *useController) SignUp(context echo.Context) error {
	user := model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := uc.uu.SignUp(user)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, res)
}

func (uc *useController) LogIn(context echo.Context) error {
	user := model.User{}

	if err := context.Bind(&user); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := uc.uu.Login(user)

	if err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	utils.SetCookie(token, context)

	return context.NoContent(http.StatusOK)
}

func (uc *useController) LogOut(context echo.Context) error {
	utils.DeleteCookie(context)

	return context.NoContent(http.StatusOK)
}

func (uc *useController) CsrfToken(context echo.Context) error {
	token := context.Get("csrf").(string)

	return context.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
