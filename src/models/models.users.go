package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FirstName string             `json:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name,omitempty" validate:"required"`
	Email     string             `json:"email,omitempty" validate:"email,required"`
	Password  string             `json:"password,omitempty" validate:"required"`
	IsAdmin   bool               `json:"isAdmin" default:"false"`
}
