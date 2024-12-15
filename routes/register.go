package routes

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"rest-api/types"
	"rest-api/utils"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func Register(c echo.Context) error {
	user := types.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	hashedPasssword := hex.EncodeToString(md5.New().Sum([]byte(user.Password)))

	stmt, _ := utils.DB.Prepare("INSERT INTO users (name, email, password) VALUES (?, ?, ?)")
	result, err := stmt.Exec(user.Name, user.Email, hashedPasssword)

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, types.Response{
			Message: "An unexpected error occured.",
		})
	}

	id, _ := result.LastInsertId()

	jwtToken := utils.CreateJWT(types.TokenUser{
		Id:    id,
		Email: user.Email,
	})

	if jwtToken == "" {
		return c.JSON(http.StatusInternalServerError, types.Response{
			Message: "An unexpected error occured.",
		})
	}

	r := types.AuthResponse{
		Message: "User created.",
		Token:   jwtToken,
	}

	return c.JSON(http.StatusOK, r)
}
