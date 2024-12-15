package routes

import (
	"fmt"
	"net/http"
	"rest-api/types"
	"rest-api/utils"

	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	id := c.Get("decoded-user-id").(int64)

	result, err := utils.DB.Query("SELECT (name) FROM users WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, types.Response{
			Message: "An unexpected error occured.",
		})
	}

	user := types.User{}
	if result.Next() {
		result.Scan(&user.Name)

	} else {
		return c.JSON(http.StatusNotFound, types.Response{
			Message: "Couldn't find user profile.",
		})
	}

	return c.JSON(http.StatusOK, types.Response{
		Message: "Welcome, " + user.Name + "!",
	})
}
