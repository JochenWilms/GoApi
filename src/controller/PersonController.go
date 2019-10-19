package controller

import (
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PersonController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"person": service.GetPerson(),
	})
}

func UpdatePerson(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"person": service.GetPerson(),
	})
}
