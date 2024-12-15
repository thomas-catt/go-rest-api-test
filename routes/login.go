package routes

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"rest-api/types"
	"rest-api/utils"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	email := c.FormValue("email")

	result, err := utils.DB.Query("SELECT * FROM users WHERE email = ? LIMIT 1", email)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, types.Response{
			Message: "An unexpected error occured.",
		})
	}

	tokenUser := types.TokenUser{}
	var user types.User
	if result.Next() {
		result.Scan(&tokenUser.Id, &user.Name, &tokenUser.Email, &user.Password)
	} else {
		return c.JSON(http.StatusNotFound, types.Response{
			Message: "Couldn't find any user with provided email.",
		})
	}

	hashedPassword := hex.EncodeToString(md5.New().Sum([]byte(c.FormValue("password"))))

	if hashedPassword != user.Password {
		return c.JSON(http.StatusForbidden, types.Response{
			Message: "Incorrect password.",
		})
	}

	jwtToken := utils.CreateJWT(tokenUser)

	if jwtToken == "" {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Message: "An unexpected error occured.",
		})
	}

	authResponse := types.AuthResponse{
		Message: "Logged in successfully.",
		Token:   jwtToken,
	}

	return c.JSON(http.StatusOK, authResponse)
}
