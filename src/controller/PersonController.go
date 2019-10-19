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
		"person": person,
	})
}
