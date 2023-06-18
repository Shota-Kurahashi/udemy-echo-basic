package utils

import (
	"ge-rest-api/src/config"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookie(token string, context echo.Context) {
	cookie := new(http.Cookie)

	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.Domain = config.GetEnv("API_DOMAIN")

	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	context.SetCookie(cookie)
}

func DeleteCookie(context echo.Context) {
	cookie := new(http.Cookie)

	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = config.GetEnv("API_DOMAIN")

	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode

	context.SetCookie(cookie)
}
