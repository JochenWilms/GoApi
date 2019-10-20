package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func SessionRouter(route *gin.Engine) {
	route.POST("/login", controller.Login)
	route.POST("/logout", controller.Logout)
}
