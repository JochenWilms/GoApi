module desyco.com/goapi

go 1.13

require (
	desyco.com/Person v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.5.0
)

replace desyco.com/Person => ./src/Person
