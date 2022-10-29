package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := SetupRouter()

	port := "8080"
	if r.Run(":"+port) != nil {
		log.Printf("Error running at port: %s", port)
	}

}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/api")

	group.POST("/test", ApiTester)

	return r
}
