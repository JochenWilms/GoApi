package Dto

import (
	"../../entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBodyDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Status string              `json:"status"`
	Id     *primitive.ObjectID `json:"id"`
	Person entity.Person       `json:"data"`
}
