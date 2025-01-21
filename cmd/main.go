package main

import (
	"fmt"
	"user/server/config"
	"user/server/controllers"
	"user/server/models"

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
	err := config.DB.AutoMigrate(&models.User{})
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
	server.GET("/technical-user", controllers.GetUser)
	server.GET("/technical-user/all", controllers.GetAllUsers)
	server.GET("/technical-user/search", controllers.SearchForUser)
	server.POST("/technical-user", controllers.CreateUser)
	server.POST("/login", controllers.Login)
}

func addProxies(server *gin.Engine) {
	err := server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Printf("Failed to set proxy %v\n", err)
		return
	}
}
