package controller

import (
	"../entity"
	"../service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPerson(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": service.GetPerson(c.Request.URL.Query()),
	})
}

func AddPerson(c *gin.Context) {
	var person entity.Person
	err := c.BindJSON(&person)
	if err != nil {
		fmt.Println(err)
	}
	service.AddPerson(person)

	c.JSON(http.StatusOK, gin.H{
		"data": person,
	})
}

type Password struct {
	Password string `json:"password"`
}

func UpdatePassword(c *gin.Context) {
	var password Password
	err := c.BindJSON(&password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password not found in body",
		})
	}
	personId, ok := c.Request.URL.Query()["id"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no person id in request.",
		})
	}

	if !service.UpdatePassword(personId[0], password.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update went not successful",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "great success",
		})
	}

}
