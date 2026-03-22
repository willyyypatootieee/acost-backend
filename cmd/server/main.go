// cmd/server/main.go
// @title Acost Backend RestAPI
// @version 1.0.0
// @description This is a simple API for managing products and orders.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"acost/docs"
	"acost/internal/router"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate go run github.com/swaggo/swag/cmd/swag@latest init -d ../../ -g cmd/server/main.go -o ../../docs --parseInternal
func main() {
	docs.SwaggerInfo.Title = "Acost Backend RestAPI"
	docs.SwaggerInfo.Description = "This is a simple API for managing products and orders."
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.RegisterRoutes(r)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
