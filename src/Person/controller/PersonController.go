package controller

import (
	"desyco.com/Person/controller/Dto"
	"desyco.com/Person/entity"
	"desyco.com/Person/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPerson(c *gin.Context) {

	person := service.GetPerson(c.Request.URL.Query())
	c.JSON(http.StatusOK, Dto.PersonDto{
		Id:   nil,
		Data: person,
	})

}

func AddPerson(c *gin.Context) {
	var person entity.Person
	err := c.BindJSON(&person)
	if err != nil {
		fmt.Println(err)
	}
	service.AddPerson(person)

	c.JSON(http.StatusOK,
		Dto.PersonDto{
			Id:   person.ID,
			Data: person,
		})
}

func UpdatePassword(c *gin.Context) {
	var password Dto.PasswordDto
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
			"error": "Update failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": "great success",
		})
	}

}
