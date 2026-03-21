// cmd/server/main.go

package main

import (
	"acost/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	router.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
