package main

import (
	"backend/model"
	"backend/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Enable CORS with all necessary headers
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3001"}, // Update with your frontend URL
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAuthorization,
			"X-Requested-With",
		},
		AllowCredentials: true,
	}))

	data, err := model.InitDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to Database")
	}
	model.DB = data

	fmt.Println("Database connection object:", data)

	routes.RegisterRoutes(e)

	// Handle preflight requests for CORS
	e.OPTIONS("/*", func(c echo.Context) error {
		return c.NoContent(204)
	})

	e.Logger.Fatal(e.Start(":1000"))
}
