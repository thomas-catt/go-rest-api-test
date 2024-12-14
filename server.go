package main

import (
	"fmt"
	"net/http"
	"rest-api/db"
	"rest-api/middlewares"
	"rest-api/routes"
	"rest-api/schema"
	"rest-api/types"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.Init()

	if db.DB == nil {
		fmt.Println("DB initialization was interrupted, exitting.")
		return
	}
	schema.Init()
	e := echo.New()
	e.Use(middlewares.Schema)

	e.POST("/register", routes.Register)
	e.GET("/login", routes.Login)
	e.GET("/users", func(c echo.Context) error {
		users := make([]types.User, 0)
		rows, err := db.DB.Query("SELECT * FROM users")

		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, types.Response{
				Message: "An unexpected error occured.",
			})
		}

		for rows.Next() {
			user := types.User{}
			var id int
			err := rows.Scan(&id, &user.Name, &user.Email, &user.Password)

			if err != nil {
				fmt.Println(err)
				return c.JSON(http.StatusInternalServerError, types.Response{
					Message: "An unexpected error occured.",
				})
			}

			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
