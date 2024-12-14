package middlewares

import (
	"fmt"
	"net/http"
	"rest-api/schema"
	"rest-api/types"

	"github.com/labstack/echo/v4"
)

func Schema(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Request().RequestURI
		fmt.Println(c.Request().Method+":\t", url)

		var routeSchema schema.RouteSchema
		schemaFound := false

		for i := range len(schema.Schema) {
			if schema.Schema[i].Route == url {
				routeSchema = schema.Schema[i]
				schemaFound = true
				break
			}
		}

		if !schemaFound {
			fmt.Println("Warning: No schema found for URL: ", url)
			return next(c)
		}

		for i := range len(routeSchema.Params) {
			param := c.FormValue(routeSchema.Params[i])
			// fmt.Println("Param: ", routeSchema.Params[i], param)
			if param == "" {
				return c.JSON(http.StatusBadRequest, types.Response{
					Message: "Missing field: " + routeSchema.Params[i],
				})
			}
		}

		return next(c)
	}
}
