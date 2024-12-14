package routes

import (
	"net/http"
	"rest-api/types"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	return c.JSON(http.StatusOK, types.Response{
		Message: "Login",
	})
}
