package router

import (
	"desyco.com/Person/controller"
	"github.com/gin-gonic/gin"
)

func PersonRouter(route *gin.Engine) {

	person := route.Group("/person")
	{
		person.GET("", controller.GetPerson)
		person.POST("", controller.AddPerson)
		person.POST("/password", controller.UpdatePassword)
	}
}
