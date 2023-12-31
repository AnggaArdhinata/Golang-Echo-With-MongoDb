package middleware

import (
	"fmt"
	"strings"

	"github.com/AnggaArdhinata/drillingfazz005/src/libs"
	"github.com/AnggaArdhinata/drillingfazz005/src/responses"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		clientToken := c.Request().Header.Get("authorization")
		if clientToken == "" {
			return c.JSON(401, responses.UserResponse{Status: 401, Message: "unauthorized", Data: &echo.Map{"data": "you must login first !"}})
		}
		if !strings.Contains(clientToken, "Bearer") {
			return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"data": "ivalid header type !"}})
		}

		tokens := strings.Replace(clientToken, "Bearer ", "", -1)

		checkToken, err := libs.ChekToken(tokens)
		if err != nil {
			return c.JSON(401, responses.UserResponse{Status: 401, Message: "unauthorized", Data: &echo.Map{"data": "unauthorized"}})
		}
		if !checkToken.IsAdmin {
			return c.JSON(401, responses.UserResponse{Status: 401, Message: "unauthorized", Data: &echo.Map{"data": "you haven't access to this feature!"}})
		}
		fmt.Println("Auth middleware pass")

		return next(c)
	}
}
