package routes

import (
	"github.com/AnggaArdhinata/drillingfazz005/src/controllers"
	"github.com/AnggaArdhinata/drillingfazz005/src/middleware"
	"github.com/labstack/echo/v4"
)

func Route() *echo.Echo {

	e := echo.New()

	api := e.Group("api/v1")

	api.GET("", func(c echo.Context) error {
		return c.String(200, "Welcome to Golang Back-End APP")

	})

	api.GET("/user", controllers.GetAllUsers)              // get all user
	api.POST("/user/register", controllers.CreateUser)     // create user/ register
	api.PUT("/user/update/:userId", controllers.EditAUser) // update user

	//apply middleware
	api.GET("/user/query", middleware.Auth(controllers.GetAUser))                // get user by id
	api.DELETE("/user/delete/:userId", middleware.Auth(controllers.DeleteAUser)) // delete by user id

	//login path
	api.POST("/auth/login", controllers.Login)

	return e
}
