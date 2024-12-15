package main

import (
	"fmt"
	"rest-api/middlewares"
	"rest-api/routes"
	"rest-api/schema"
	"rest-api/utils"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	utils.InitDB()

	if utils.DB == nil {
		fmt.Println("DB initialization was interrupted, exitting.")
		return
	}
	schema.Init()
	e := echo.New()
	e.Use(middlewares.Schema)

	e.POST("/register", routes.Register)
	e.POST("/login", routes.Login)
	e.GET("/profile", routes.Profile, middlewares.Auth)

	e.Logger.Fatal(e.Start(":3000"))
}
