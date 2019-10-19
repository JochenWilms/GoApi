package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func PersonRouter(route *gin.Engine) {
	person := route.Group("/person")
	{
		person.GET("", controller.PersonController)
	}
}
