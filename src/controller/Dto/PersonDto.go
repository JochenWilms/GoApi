package Dto

import (
	"../../entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonDto struct {
	Id   *primitive.ObjectID `json:"id"`
	Data entity.Person       `json:"data"`
}

type PasswordDto struct {
	Password string `json:"password"`
}

type PersonErrorDto struct {
	Error string `json:"error"`
}
