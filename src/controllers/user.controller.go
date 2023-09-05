package controllers

import (
	"context"
	"time"

	"github.com/AnggaArdhinata/drillingfazz005/src/configs"
	"github.com/AnggaArdhinata/drillingfazz005/src/libs"
	"github.com/AnggaArdhinata/drillingfazz005/src/models"
	"github.com/AnggaArdhinata/drillingfazz005/src/responses"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, responses.UserResponse{Status: 400, Message: "Bad Request", Data: &echo.Map{"users": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(400, responses.UserResponse{Status: 400, Message: "Bad Request", Data: &echo.Map{"users": validationErr.Error()}})
	}

	newUser := models.User{
		Id:        primitive.NewObjectID(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		IsAdmin:   user.IsAdmin,
	}

	hashPass, err := libs.HashPassword(newUser.Password)
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err}})
	}
	newUser.Password = hashPass
	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
	}

	return c.JSON(201, responses.UserResponse{Status: 201, Message: "created", Data: &echo.Map{"users": result}})
}

func GetAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.QueryParam("id")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
	}
	user.Password = "scrett!"
	return c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: &echo.Map{"users": user}})
}

func EditAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, responses.UserResponse{Status: 400, Message: "bad request", Data: &echo.Map{"users": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(400, responses.UserResponse{Status: 400, Message: "bad request", Data: &echo.Map{"users": validationErr.Error()}})
	}

	update := bson.M{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "password": user.Password}
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "bad request", Data: &echo.Map{"users": err.Error()}})
	}

	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
		}
	}

	return c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: &echo.Map{"users": updatedUser}})
}

func DeleteAUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(404, responses.UserResponse{Status: 404, Message: "not found", Data: &echo.Map{"users": "User with specified ID not found!"}})
	}

	return c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: &echo.Map{"users": "User successfully deleted!"}})
}

func GetAllUsers(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.JSON(500, responses.UserResponse{Status: 500, Message: "error", Data: &echo.Map{"users": err.Error()}})
		}
		singleUser.Password = "scrett!"
		users = append(users, singleUser)
	}

	return c.JSON(200, responses.UserResponse{Status: 200, Message: "success", Data: &echo.Map{"users": users}})
}
