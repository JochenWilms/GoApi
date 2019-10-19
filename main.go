package main

import (
	"./src/controller"
	"./src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controller.IndexController)

	router.PersonRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
