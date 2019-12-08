package main

import (
	"desyco.com/Person/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.PersonRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
