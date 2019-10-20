package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"

	"../service"
	"./Dto"
)

func Login(c *gin.Context) {
	fmt.Println("login")
	var body Dto.LoginBodyDto
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println(err)
	}

	person := service.FindPersonByMail(body.Email)

	session := sessions.Default(c)

	if person.VerifyPassword(body.Password) {
		session.Set("login", person.ID.Hex())
		session.Save()
		c.JSON(http.StatusOK, Dto.LoginResponseDto{
			Status: "Login succeed",
			Id:     person.ID,
			Person: person,
		})
	} else {
		session.Set("login", nil)
		session.Save()
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password",
		})
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)

	fmt.Println(session.Get("login"))

	session.Set("login", nil)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"data": "Logout succeeded",
	})
}
