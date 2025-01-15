package main

import (
	"example/server/config"
	"example/server/controllers"
	"example/server/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	setupServer(server)
	setupDataBase()
	err := server.Run("localhost:8080")
	if err != nil {
		return
	}
}

func setupDataBase() {
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.TechnicalUser{})
	if err != nil {
		fmt.Printf("Failed to create tables")
	}
}

func setupServer(server *gin.Engine) {
	addCorsConfig(server)
	addRouters(server)
	addProxies(server)
}

func addCorsConfig(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	server.Use(cors.New(corsConfig))
}

func addRouters(server *gin.Engine) {
	server.GET("/technical-user", controllers.GetTechnicalUser)
	server.GET("/technical-user/search", controllers.SearchForTechnicalUser)
	server.POST("/technical-user", controllers.CreateTechnicalUser)
}

func addProxies(server *gin.Engine) {
	err := server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Printf("Failed to set proxy %v\n", err)
		return
	}
}
