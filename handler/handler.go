package handler

import (
	"golang-graphql-authentication/auth"
	"net/http"

	"github.com/labstack/echo"
)

func Welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome")
	}
}

type Token struct {
	Token string `json:"token"`
}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "you" && password == "passed" {
			token, err := auth.GenerateToken()

			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, Token{
				Token: token,
			})
		}

		return echo.ErrUnauthorized
	}
}
