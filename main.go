package main

import (
	"./src/controller"
	"./src/router"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	router.SessionRouter(r)
	router.PersonRouter(r)

	r.GET("/", controller.IndexController)

	r.Run() // listen and serve on 0.0.0.0:8080
}
