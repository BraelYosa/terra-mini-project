package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.Login)

	userGroup := e.Group("/auth", middleware.CheckToken)
	userGroup.GET("/protected", controllers.Protected)

	dataGroup := e.Group("/data", middleware.CheckToken)

	dataGroup.GET("", controllers.GetData)
	dataGroup.POST("", controllers.SubmitData)
}
