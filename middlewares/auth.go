package middlewares

import (
	"net/http"
	"rest-api/types"
	"rest-api/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenUser, ok := utils.DecodeJWT(strings.ReplaceAll(authHeader, "Bearer ", ""))

		if !ok {
			return c.JSON(http.StatusUnauthorized, types.Response{
				Message: "Session expired",
			})
		}

		c.Set("decoded-user-id", tokenUser.Id)
		return next(c)
	}
}
