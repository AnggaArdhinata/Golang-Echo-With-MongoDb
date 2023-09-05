package routes

import (
	"github.com/AnggaArdhinata/drillingfazz005/src/controllers"
	"github.com/AnggaArdhinata/drillingfazz005/src/middleware"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {
	
	e := echo.New()

	e.POST("/user/register", controllers.CreateUser)
	e.GET("/user/query", controllers.GetAUser)
	e.PUT("/user/:userId", controllers.EditAUser)
	e.DELETE("/user/:userId", middleware.Auth(controllers.DeleteAUser))
	e.GET("/user", controllers.GetAllUsers)

	e.POST("/auth/login", controllers.Login)

	return e
}