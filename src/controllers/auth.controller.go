package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/AnggaArdhinata/drillingfazz005/src/libs"
	"github.com/AnggaArdhinata/drillingfazz005/src/models"
	"github.com/AnggaArdhinata/drillingfazz005/src/responses"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)


func Login(c echo.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	email := c.FormValue("email")
	password := c.FormValue("password")
	var user models.User
	defer cancel()

	// validate request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: &echo.Map{"data": err.Error()}})
	}

	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "InternalServerError", Data: &echo.Map{"message": "wrong email or password !"}})
	}

	passwordIsValid := libs.CheckPassword(user.Password, password)
	if !passwordIsValid {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "InternalServerError", Data: &echo.Map{"message": "wrong email or password !"}})
	}
	
	jwt := libs.NewToken(user.Id, user.IsAdmin)
	token, err := jwt.Create()
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "InternalServerError", Data: &echo.Map{"message": err.Error()}})
	}
	return c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: &echo.Map{"token": token}})
}
